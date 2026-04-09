package voting

import (
	"crypto/rand"
	"math/big"

	"golang.org/x/crypto/sha3"
)

// CommitData содержит данные для фазы commit.
// Salt необходимо сохранить — он нужен для reveal!
type CommitData struct {
	CandidateID *big.Int
	Salt        []byte   // 32 байта, хранить у пользователя
	Hash        [32]byte // отправить в commit()
}

// NewCommit генерирует случайный salt и вычисляет commit hash.
// Воспроизводит Solidity: keccak256(abi.encodePacked(candidateId, salt))
func NewCommit(candidateID *big.Int) (*CommitData, error) {
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	// encodePacked для uint256: дополняем до 32 байт
	idBytes := make([]byte, 32)
	candidateID.FillBytes(idBytes)

	h := sha3.NewLegacyKeccak256()
	h.Write(idBytes)
	h.Write(salt)

	var hash [32]byte
	copy(hash[:], h.Sum(nil))

	return &CommitData{
		CandidateID: candidateID,
		Salt:        salt,
		Hash:        hash,
	}, nil
}