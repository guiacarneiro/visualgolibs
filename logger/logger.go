/*
 * Created on: 15/01/2019
 *     Author: guilhermehenrique
*/

package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	Debug   = 4
	Info    = 3
	Warning = 2
	Error   = 1
	Fatal   = -1
)

var (
	onceLogger        sync.Once
	mutexLogger       = &sync.Mutex{}
	caminhoArquivoLog string
	nomeArquivoLog 	  string
	logLevel          int
	versao 			  int

	//LogFatal - log do tipo Fatal
	LogFatal = log.New(logWriter{Fatal}, "FATAL: ", 0)

	//LogError - log do tipo Error
	LogError = log.New(logWriter{Error}, "ERROR: ", 0)

	//LogWarning - log do tipo Warning
	LogWarning = log.New(logWriter{Warning}, "WARN: ", 0)

	//LogInfo - log do tipo Info
	LogInfo = log.New(logWriter{Info}, "INFO: ", 0)

	//LogDebug - log do tipo Debug
	LogDebug = log.New(logWriter{Debug}, "DEBUG: ", 0)
)

func init() {
}

func Inicializa(level int, logPath string, logFileName string, projectVersion int) {
	onceLogger.Do(func() {
		logLevel = level
		caminhoArquivoLog = logPath
		nomeArquivoLog = logFileName
		versao = projectVersion
		var _, err = os.Stat(caminhoArquivoLog)
		if os.IsNotExist(err) {
			err = os.MkdirAll(caminhoArquivoLog, os.ModePerm)
			if err != nil {
				panic(fmt.Errorf("erro ao criar diretorio de log: %s", err))
			}
		}
	})
}

type logWriter struct {
	logLevel int
}

func (f logWriter) Write(p []byte) (n int, err error) {
	pc, file, line, ok := runtime.Caller(4)
	if !ok {
		file = "?"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	gravaLog(Error, fmt.Sprintf("{%d} %s:%d %s: %s", versao, filepath.Base(file), line, fnName, p))
	return len(p), nil
}

func gravaLog(tipoLog int, mensagem string) {
	if logLevel == 0 {
		panic("logLevel nao informado")
	}
	if caminhoArquivoLog == "" {
		panic("caminho dos arquivos logs nao informado")
	}
	if versao == 0 {
		panic("versao do projeto nao informada")
	}

	mutexLogger.Lock()
	defer mutexLogger.Unlock()

	if tipoLog > logLevel {
		return
	}

	current := time.Now()

	var log = strings.Join([]string{caminhoArquivoLog, "/", nomeArquivoLog, "_", strconv.Itoa(current.Year()),
		fmt.Sprintf("%02d", current.Month()),
		fmt.Sprintf("%02d", current.Day()), ".log"}, "")

	var file, err = os.OpenFile(log, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(fmt.Errorf("erro ao abrir arquivo de log: %s", err))
	}
	defer file.Close()

	_, err = file.WriteString(time.Now().Format("[02/01/2006 - 15:04:05]") + mensagem)
	if err != nil {
		panic(fmt.Errorf("erro ao escrever no arquivo de log: %s", err))
	}

	err = file.Sync()
	if err != nil {
		panic(fmt.Errorf("erro ao salvar arquivo de log: %s", err))
	}

	if tipoLog == Fatal {
		panic(mensagem)
	}
}
