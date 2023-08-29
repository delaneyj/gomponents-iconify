package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m5.962 20.174l-.383.384a1.51 1.51 0 1 1-2.137-2.137l.384-.383a2.014 2.014 0 0 0 .578-1.647l-.09-.799a3.021 3.021 0 0 1 .867-2.47l8.943-8.942a4.028 4.028 0 1 1 5.696 5.696l-8.942 8.943a3.02 3.02 0 0 1-2.47.866l-.8-.089a2.014 2.014 0 0 0-1.646.578Z" opacity=".5"/><path d="m12.518 17.18l-1.061 1.06l-1.3-1.301a.75.75 0 1 1 1.06-1.06l1.3 1.3Zm3.209-3.21l-1.06 1.06l-1.301-1.3a.75.75 0 0 1 1.06-1.06l1.3 1.3Zm3.199-3.2l-1.06 1.061l-1.301-1.3a.75.75 0 1 1 1.06-1.061l1.301 1.3Z"/><path fill-rule="evenodd" d="M16.03 7.97a.75.75 0 0 1 0 1.06l-6.5 6.5a.75.75 0 0 1-1.06-1.06l6.5-6.5a.75.75 0 0 1 1.06 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}