package TEST

import (
	"fmt"
	"sync"
	"time"
)

// 演示 sync.Once 的工作原理

var (
	configLoaded bool
	once         sync.Once
)

type Config struct {
	DatabaseURL string
	APIKey      string
}

var appConfig Config

// 模拟加载配置的耗时操作
func loadConfig() Config {
	fmt.Println("正在加载配置...")
	time.Sleep(2 * time.Second) // 模拟耗时操作
	return Config{
		DatabaseURL: "mysql://localhost:3306/mydb",
		APIKey:      "secret-key-12345",
	}
}

// 使用 sync.Once 确保配置只加载一次
func GetConfigWithOnce() Config {
	once.Do(func() {
		fmt.Println("=== 第一次调用，执行加载操作 ===")
		appConfig = loadConfig()
		configLoaded = true
		fmt.Printf("配置加载完成: %+v\n", appConfig)
	})
	return appConfig
}

// 不使用 sync.Once 的版本（有问题的版本）
func GetConfigWithoutOnce() Config {
	if !configLoaded {
		fmt.Println("=== 检查到配置未加载，执行加载操作 ===")
		appConfig = loadConfig()
		configLoaded = true
		fmt.Printf("配置加载完成: %+v\n", appConfig)
	}
	return appConfig
}

func testOnce() {
	fmt.Println("=== 演示 sync.Once 的作用 ===")
	fmt.Println()

	// 重置状态
	configLoaded = false
	once = sync.Once{}
	appConfig = Config{}

	fmt.Println("1. 使用 sync.Once 的情况:")
	fmt.Println("----------------------------------------")

	// 模拟多个goroutine同时访问配置
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			config := GetConfigWithOnce()
			fmt.Printf("Goroutine %d 获取到配置: %+v\n", id, config)
		}(i)
	}
	wg.Wait()

	fmt.Println()
	fmt.Println("2. 不使用 sync.Once 的情况 (有并发问题):")
	fmt.Println("----------------------------------------")

	// 重置状态
	configLoaded = false
	appConfig = Config{}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			config := GetConfigWithoutOnce()
			fmt.Printf("Goroutine %d 获取到配置: %+v\n", id, config)
		}(i)
	}
	wg.Wait()

	fmt.Println()
	fmt.Println("=== sync.Once 的关键特性 ===")
	fmt.Println("1. 线程安全: 确保在并发环境下只执行一次")
	fmt.Println("2. 原子性操作: 使用底层原子操作保证线程安全")
	fmt.Println("3. 内存屏障: 确保执行完成后的结果对所有goroutine可见")
	fmt.Println("4. 简单易用: 只需调用 once.Do(func())")
}
