# go_CLI_counter

🧾 项目名称：Go Word Frequency Counter
📌 项目目标

构建一个命令行工具，读取文本文件，统计每个单词出现的次数，并按频率排序输出。要求支持多个输入文件，具备基本的可扩展性与错误处理能力。
✅ 功能需求
输入

    支持通过命令行参数传入一个或多个文本文件路径（如：go run main.go file1.txt file2.txt）

    支持选择是否区分大小写（flag：--case-sensitive）

输出

    控制台输出单词频率统计结果，格式如：

    the: 123
    is: 97
    golang: 56

    支持按频率排序或按字典序排序（flag：--sort=freq|alpha）

扩展功能（可选）

    忽略停用词（如“the”, “is”, “a”等常见词）

    支持输出为 JSON 或写入文件（flag：--out=json|text）

🏗️ 项目结构建议（初版）

```
wordfreq/
├── main.go                 // 入口函数，处理 CLI 参数
├── reader/
│   └── reader.go           // 文件读取相关逻辑
├── counter/
│   └── counter.go          // 词频统计逻辑
├── printer/
│   └── printer.go          // 排序与输出逻辑
├── utils/
│   └── text_cleaner.go     // 文本清洗（去标点、小写化）
├── go.mod                  // Go 模块文件
└── test_data/
    └── sample.txt          // 测试用文本文件

