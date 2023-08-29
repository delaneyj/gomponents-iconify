package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func File(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="#F1C40F" d="M0 100V0h72l28 28v72H0z"/><path fill="#F39C12" d="M0 100V72L72 0l28 28l-72 72H0z"/><path fill="#FFF55B" d="M72 0h1l27 27v1H72V0z"/><path fill="#fff" d="M34.828 55.721v3.853H23.29v7.704h9.738v3.87H23.29V81.75h-4.878V55.721h16.416zm9.252 26.028h-4.86V55.721h4.86v26.028zm10.691-3.996h10.422v3.996H49.93V55.721h4.842v22.032zm30.187-22.032v3.853H73.42v7.218h9.09v3.726h-9.09v7.362h11.538v3.87H68.542V55.721h16.416z"/>`),
		g.Group(children),
	)
}