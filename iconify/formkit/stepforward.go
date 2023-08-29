package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stepforward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2.5 12.7c-.25 0-.5-.06-.73-.19A1.49 1.49 0 0 1 1 11.2V4.8c0-.55.29-1.04.77-1.31c.48-.27 1.05-.25 1.52.04l5.11 3.2c.44.28.71.75.71 1.27s-.26 1-.71 1.27l-5.11 3.2c-.25.15-.52.23-.79.23Zm0-8.4c-.11 0-.2.04-.25.06c-.08.04-.26.17-.26.44v6.39c0 .27.18.39.26.44c.08.04.28.13.51-.01l5.11-3.2c.21-.13.24-.34.24-.42s-.02-.29-.24-.42l-5.11-3.2a.461.461 0 0 0-.26-.08Zm9 8.2c-.28 0-.5-.22-.5-.5V4c0-.28.22-.5.5-.5s.5.22.5.5v8c0 .28-.22.5-.5.5Z"/>`),
		g.Group(children),
	)
}