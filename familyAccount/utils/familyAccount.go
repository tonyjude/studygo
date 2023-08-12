package utils

import "fmt"

type familyAccount struct {
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	details string
}

func (family *familyAccount) detail() {
	fmt.Println("------------------------当前收支明细记录------------------------")
	fmt.Println(family.details)
}

func (family *familyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&family.money)
	family.balance += family.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&family.note)
	family.details += fmt.Sprintf("\n收入 \t %v \t\t %v \t\t %v", family.balance, family.money, family.note)
}

func (family *familyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&family.money)
	if family.money > family.balance {
		fmt.Println("余额不足！")
	}
	family.balance -= family.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&family.note)
	family.details += fmt.Sprintf("\n支出 \t %v \t\t %v \t\t %v", family.balance, family.money, family.note)
}

func (family *familyAccount) exit() {
	fmt.Println("你确定要退出吗？ y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}

	if choice == "y" {
		family.loop = false
	}
}

func NewFamilyAccount() *familyAccount {
	return &familyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		details: "收支 \t 账户金额 \t 收支金额 \t 说   明",
	}
}

func (family *familyAccount) MainMenu() {
	for {
		fmt.Println("\n------------------------家庭收支记账软件------------------------")
		fmt.Println("                         1 收支明细")
		fmt.Println("                         2 登记收入")
		fmt.Println("                         3 登记支出")
		fmt.Println("                         4 退出软件")
		fmt.Print("请选择（1-4）：")

		fmt.Scanln(&family.key)

		switch family.key {
		case "1":
			family.detail()
		case "2":
			family.income()
		case "3":
			family.pay()
		case "4":
			family.exit()
		default:
			fmt.Println("请输入正确的选项。。。")
		}

		if !family.loop {
			break
		}
	}

	fmt.Println("退出成功")
}
