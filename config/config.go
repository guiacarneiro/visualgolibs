package config

import (
	"fmt"
	"sync"
	"visualgolibs/logger"

	"github.com/spf13/viper"
)

var (
	//Versao - versao da aplicacao
	Versao                  = 0
	nomeArquivoConfiguracao string
	once                    sync.Once
)

//Inicializa - inicializa configuracoes
func Inicializa(nomeArquivoConfig string, versao int) {
	nomeArquivoConfiguracao = nomeArquivoConfig
	Versao = versao
}

//GetPropriedade - busca propriedade no arquivo de configuracoes
func GetPropriedade(propriedade string) string {
	once.Do(func() {
		viper.SetConfigName(nomeArquivoConfiguracao)
		viper.AddConfigPath("./")
		viper.AddConfigPath("./config/")

		errConfig := viper.ReadInConfig()
		if errConfig != nil {
			logger.LogError.Println(fmt.Errorf("erro ao buscar arquivo de configurações: %s", errConfig))
			panic(fmt.Errorf("erro ao buscar arquivo de configurações: %s", errConfig))
		}
	})
	return viper.GetString(propriedade)
}

//GetPropriedadeInt - busca propriedade no arquivo de configuracoes
func GetPropriedadeInt(propriedade string) int {
	once.Do(func() {
		viper.SetConfigName(nomeArquivoConfiguracao)
		viper.AddConfigPath("./")
		viper.AddConfigPath("./config/")

		errConfig := viper.ReadInConfig()
		if errConfig != nil {
			logger.LogError.Println(fmt.Errorf("erro ao buscar arquivo de configurações: %s", errConfig))
			panic(fmt.Errorf("erro ao buscar arquivo de configurações: %s", errConfig))
		}
	})
	return viper.GetInt(propriedade)
}

//UnmarshalProperties - busca propriedade no arquivo de configuracoes
func UnmarshalProperties(obj interface{}) error {
	once.Do(func() {
		viper.SetConfigName(nomeArquivoConfiguracao)
		viper.AddConfigPath("./")
		viper.AddConfigPath("./config/")

		errConfig := viper.ReadInConfig()
		if errConfig != nil {
			logger.LogError.Println(fmt.Errorf("erro ao buscar arquivo de configurações: %s", errConfig))
			panic(fmt.Errorf("erro ao buscar arquivo de configurações: %s", errConfig))
		}
	})
	err := viper.Unmarshal(obj)
	return err
}