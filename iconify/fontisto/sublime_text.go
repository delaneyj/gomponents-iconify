package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SublimeText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M.003 23.617V17.93a.72.72 0 0 1 .457-.654l.005-.002l7.453-2.361l-7.454-2.366a.706.706 0 0 1-.398-.377l-.002-.005a.434.434 0 0 1-.061-.222v-.014v.001v-5.737c0-.083.023-.161.064-.227l-.001.002a.7.7 0 0 1 .395-.379l.005-.002L18.014.023a.343.343 0 0 1 .461.365V.386v5.686a.724.724 0 0 1-.457.654l-.005.002l-7.375 2.338l7.378 2.339a.722.722 0 0 1 .462.656v5.691a.53.53 0 0 1-.011.106l.001-.003a.725.725 0 0 1-.45.558l-.005.002L.464 23.979a.429.429 0 0 1-.128.022a.345.345 0 0 1-.332-.387v.002z"/>`),
		g.Group(children),
	)
}