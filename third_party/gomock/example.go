//go:generate mockgen -destination mock_example.go -package gomock study/leetcode/third_party/gomock Repository
package gomock

type Repository interface {
	Create(key string, value []byte) error
	Retrieve(key string) ([]byte, error)
	Update(key string, value []byte) error
	Delete(key string) error
}
