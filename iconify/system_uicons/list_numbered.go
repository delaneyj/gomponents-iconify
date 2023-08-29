package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListNumbered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.5 10.5h-7m7 4h-7m7-8h-7"/><path fill="currentColor" d="M5.88 8V5.828h-.037l-.68.459V5.67l.717-.488h.717V8zm-.98 2.068c0-.572.45-.963 1.109-.963c.652 0 1.04.354 1.04.836c0 .334-.148.555-.597.961l-.555.502v.037h1.186V12H4.94v-.479l1.008-.912c.348-.318.406-.44.406-.605c0-.195-.136-.358-.382-.358c-.262 0-.416.178-.416.422zm.712 4.73v-.484h.362c.238 0 .392-.138.392-.341c0-.192-.146-.332-.388-.332c-.254 0-.409.134-.42.363h-.653c.01-.541.438-.899 1.108-.899c.66 0 1.021.346 1.02.766c0 .34-.22.565-.528.637v.037c.406.057.64.309.64.68c0 .504-.48.851-1.158.851c-.67 0-1.125-.361-1.15-.916h.684c.01.217.185.352.457.352c.261 0 .439-.143.439-.356c0-.222-.168-.357-.443-.357z"/></g>`),
		g.Group(children),
	)
}