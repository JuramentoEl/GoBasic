package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	env "github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type (
	Analyzer struct {
		File   string `env:"FILE"`
		Level  string `env:"LEVEL"`
		Output string `env:"OUTPUT"`
	}
	Cfg struct {
		Analyzer Analyzer `envPrefix:"ANALYZER_"`
	}
)

func main() {
	var pathFile, levelLog, pathOutput string
	var count int

	flag.StringVar(&pathFile, "file", "", "Путь к файлу")
	flag.StringVar(&levelLog, "level", "", "Уровень логов")
	flag.StringVar(&pathOutput, "output", "", "Путь к файлу статистики")
	flag.Parse()

	err := defaultData(&pathFile, &levelLog, &pathOutput)
	if err != nil {
		fmt.Println(err)
	}

	lines, err := readCsv(pathFile)
	if err != nil {
		fmt.Println(err)
	}

	data := make(map[string]int)

	for _, line := range lines {
		if levelLog == "info" || (levelLog == "warning" && line[1] != "info") || (levelLog == "error" && line[1] == "error") {
			count++
			data[line[2]]++
		}
	}
	writeCsv(pathOutput, count, data)
	fmt.Println(count)
}

func defaultData(pathFile, levelLog, pathOutput *string) error {
	if *pathFile != "" && *levelLog != "" && *pathOutput != "" {
		return nil
	}

	err := load(os.Getenv("ENV"))
	if err != nil {
		log.Fatal("load config failed:", err)
	}
	cfg, err := initCfg()
	if err != nil {
		return err
	}

	if *pathFile == "" {
		*pathFile = cfg.Analyzer.File
	}
	if *levelLog == "" {
		*levelLog = cfg.Analyzer.Level
	}
	if *pathOutput == "" {
		*pathOutput = cfg.Analyzer.Output
	}
	return nil
}

func load(cfgFile string) error {
	if len(cfgFile) == 0 {
		cfgFile = ".env"
	}

	return godotenv.Load(cfgFile)
}

func initCfg() (*Cfg, error) {
	cfg := &Cfg{}
	opts := env.Options{
		Prefix:                "LOG_",
		UseFieldNameByDefault: true,
	}

	if err := env.ParseWithOptions(cfg, opts); err != nil {
		return nil, err
	}
	return cfg, nil
}

func readCsv(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open the CSV file")
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read the CSV file")
	}

	return lines, nil
}

func writeCsv(path string, count int, data map[string]int) error {
	var s string
	for i, element := range data {
		s = s + i + ": " + strconv.Itoa(element) + "\n"
	}
	s = s + "\n" + "Количество записей: " + strconv.Itoa(count)

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(s))
	if err != nil {
		return err
	}

	return nil
}
