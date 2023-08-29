package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinimalisticMagniferZoomOutBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.064 20.286c5.006 0 9.064-4.094 9.064-9.143C20.128 6.093 16.07 2 11.064 2S2 6.093 2 11.143c0 5.05 4.058 9.143 9.064 9.143Zm8.657-2.503c-1.085-.09-1.99.823-1.902 1.918c.016.19.084.416.153.648l.02.067l.018.06c.063.213.125.424.21.585a1.754 1.754 0 0 0 2.527.643c.151-.1.305-.257.46-.414l.045-.045l.044-.044a3.31 3.31 0 0 0 .41-.465a1.786 1.786 0 0 0-.637-2.549a3.254 3.254 0 0 0-.58-.212l-.06-.018l-.066-.02c-.23-.07-.454-.138-.642-.154ZM9.05 10.381a.759.759 0 0 0-.756.762c0 .42.339.762.756.762h4.028a.759.759 0 0 0 .756-.762a.759.759 0 0 0-.756-.762H9.05Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}