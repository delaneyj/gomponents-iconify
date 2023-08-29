package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnhealthyFoodNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsUnhealthyFoodNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm15.23 5.807L17 12.41V8.79l7-2.14v3.893l6.086-2.913l-.069 1.01L31.227 4h7.117l-2.769 9.23l1.767-.452l-2.148 8.014c.14-.135.28-.271.421-.41a1.293 1.293 0 0 1 2.203 1.071L35.22 42.248A2 2 0 0 1 33.234 44H14.766a2 2 0 0 1-1.985-1.752l-2.599-20.79a1.297 1.297 0 0 1 1.886-1.31l-2.44-9.104l5.602-5.237ZM17 24.293a31.17 31.17 0 0 1-2.214-1.73L11.87 11.686l2.29-2.14L17 20.131v4.161Zm5 2.428c-.979-.23-1.973-.623-3-1.177V10.269l3-.917V26.72Zm4.797-.185a9.216 9.216 0 0 1-2.535.426l1.006-14.81l2.59-1.24l-1.061 15.624Zm5.675-3.31c-1.116.909-2.178 1.65-3.209 2.228l2.328-8.687l2.84-.852l-1.96 7.312Zm.842-9.418L35.656 6h-2.883l-2.221 8.514l2.762-.706Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsUnhealthyFoodNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}