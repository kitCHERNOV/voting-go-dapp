// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract VoterRegistry is AccessControl {
    bytes32 public constant REGISTRAR_ROLE = keccak256("REGISTRAR_ROLE");

    mapping(address => bool) public isRegistered;
    bool public selfRegistrationOpen;

    event VoterRegistered(address indexed voter, address indexed registrar);
    event VoterRevoked(address indexed voter, address indexed registrar);

    constructor(address[] memory registrars) {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        for (uint256 i = 0; i < registrars.length; i++) {
            _grantRole(REGISTRAR_ROLE, registrars[i]);
        }
    }

    /// @notice Регистрирует участника. Только REGISTRAR_ROLE.
    function register(address _voter) external onlyRole(REGISTRAR_ROLE) {
        require(!isRegistered[_voter], "Already registered");
        isRegistered[_voter] = true;
        emit VoterRegistered(_voter, msg.sender);
    }

    /// @notice Отзывает регистрацию участника. Только REGISTRAR_ROLE.
    function revoke(address _voter) external onlyRole(REGISTRAR_ROLE) {
        require(isRegistered[_voter], "Not registered");
        isRegistered[_voter] = false;
        emit VoterRevoked(_voter, msg.sender);
    }

    /// @notice Самостоятельная регистрация (если флаг открыт).
    function selfRegister() external {
        require(selfRegistrationOpen, "Self-registration closed");
        require(!isRegistered[msg.sender], "Already registered");
        isRegistered[msg.sender] = true;
        emit VoterRegistered(msg.sender, msg.sender);
    }

    /// @notice Открыть/закрыть самостоятельную регистрацию. Только DEFAULT_ADMIN_ROLE.
    function setSelfRegistration(bool _open) external onlyRole(DEFAULT_ADMIN_ROLE) {
        selfRegistrationOpen = _open;
    }
}