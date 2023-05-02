package levelDB

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/yin-zt/itsm-workflow/pkg/config"
	"github.com/yin-zt/itsm-workflow/pkg/utils/loger"
)

var (
	CONST_LEVELDB_FILE_NAME = config.DATA_DIR + "cmdb.db"
	opts                    = &opt.Options{
		CompactionTableSize: 1024 * 1024 * 20,
		WriteBuffer:         1024 * 1024 * 20,
	}
	logger = loger.GetCmdbLogger()
	//UserRelationField = make(map[string]int)
)

type LevelDb struct {
	DB *leveldb.DB
}

func NewLevelDb() *LevelDb {
	db, err := leveldb.OpenFile(CONST_LEVELDB_FILE_NAME, opts)
	if err != nil {
		logger.Errorf("NewLevelDb error, with: %v", err)
	}
	return &LevelDb{db}
}

func (this LevelDb) GetKeyFromDB(key string) (string, error) {
	if data, err := this.DB.Get([]byte(key), nil); err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}

func (this LevelDb) SetKeyInLevelDB(key string, value interface{}) error {
	switch trueVal := value.(type) {
	case string:
		err := this.DB.Put([]byte(key), []byte(trueVal), nil)
		return err
	case []string, map[string]interface{}, []map[string]string:
		valStr, err := json.Marshal(trueVal)
		if err != nil {
			logger.Errorf("[]string格式数据序列化失败, 值为: %v", trueVal)
			return err
		}
		err = this.DB.Put([]byte(key), valStr, nil)
		return err
	default:
		logger.Errorf("Unknow the type of value, value is: %v", trueVal)
		return nil
	}
}
