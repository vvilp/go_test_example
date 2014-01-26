package main 

import "code.google.com/p/go.net/html"
import "fmt"
import "strings"



func traverse_html_node(node *html.Node, deep int) {
	space := ""
	for i := 0 ; i < deep ; i++ {
		space += "\t"
	}

	fmt.Println(space, "node.Type : ", node.Type)
	fmt.Println(space, "node.Data : ", node.Data)
	fmt.Println(space, "node.Attr : ", node.Attr)
	fmt.Println(space, "node.Namespace : ", node.Namespace)

	for n := node.FirstChild ; n != nil ; n = n.NextSibling {
		traverse_html_node(n, deep + 1)
	}
}

func traverse_html_tokenizer(z *html.Tokenizer) {
	for {
		if z.Next() == html.ErrorToken {
			return
		}
		text_b := z.Text()
		tag_name_b, hasAttri := z.TagName()
		tag_attr_key_b, tag_attr_value_b, _ := z.TagAttr()
		text 			:= string(text_b)
		tag_name 		:= string(tag_name_b)
		tag_attr_key 	:= string(tag_attr_key_b)
		tag_attr_value 	:= string(tag_attr_value_b)
		fmt.Printf("|Tokenizer.Text:%-10s|Tokenizer.TagName:%-10s|hasAttri:%-10t|tag_attr_key:%-10s|tag_attr_value:%-10s|\n", text, tag_name, hasAttri,tag_attr_key,tag_attr_value)
	}
}

func traverse_html_token(z *html.Tokenizer) {
	for {
		if z.Next() == html.ErrorToken {
			return
		}
		token := z.Token()
		token_type := token.Type
		fmt.Printf("|token_type:%-20s|token.Data:%-10s|token.Attr:%-10s|\n",token_type,token.Data,token.Attr)
	}
}

func main() {
	s := `<p>Links:<a href="a1" class="test"/></p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	
	doc, _ := html.Parse(strings.NewReader(s))
	traverse_html_node(doc, 0)

	z := html.NewTokenizer(strings.NewReader(s))
	traverse_html_tokenizer(z)

	z1 := html.NewTokenizer(strings.NewReader(s))
	traverse_html_token(z1)
}