package main

import (
	"log"
	"unicode"
)

var (
	ruler map[string]string
)

func (s *Stack) Solve() (text string) { //字符串处理
	if s.Top() == ")" { //成功匹配出一套字符串
		s.Pop()
		for s.Top() != "(" {
			if s.Empty() {
				text = "括号匹配失败"
				return text
			}
			text = s.Top() + text
			s.Pop()
		}
		s.Pop()

		l := len(text) - 1
		if l < 0 {
			text = "输入非法"
			return text
		}
		//处理第二类产生式
		for i := l; i > 0; i-- {
			s.Push(string(text[0]))
			s.Push(string(text[i]))
		}
		s.Push(string(text[0]))
	}
	if unicode.IsUpper(rune(s.Top()[0])) { //是大写字母，对应字符串入栈
		text = ruler[s.Top()]
		if text == "" {
			text = "输入的 " + s.Top() + " 未指定对应字符串"
			return text
		}
		s.Pop()
		for _, v := range text {
			s.Push(string(v))
		}
	}
	return ""
}
func Translator(input string) (output string) {
	stack := NewStack()
	for _, v := range input {
		if unicode.IsLetter(v) { // 是字母
			stack.Push(string(v))
			err := stack.Solve()
			if err != "" {
				output = err
				log.Println(output)
				return output
			}
		} else if v == '(' || v == ')' { //不是字母，判断是否为"("，否则非法
			stack.Push(string(v))
			err := stack.Solve()
			if err != "" {
				output = err
				log.Println(output)
				return output
			}
		} else {
			output = "错误，输入非法"
			log.Println(output)
			return output
		}

	}
	for !stack.Empty() { //开始输出字符
		err := stack.Solve()
		if err != "" {
			output = err
			log.Println(output)
			return output
		}
		v := stack.Top()
		stack.Pop()
		if v == "(" || v == ")" { //存在非字母字符
			output = "错误，括号匹配未完整"
			log.Println(output)
			return output
		}
		output = v + output
	}
	return output
}
