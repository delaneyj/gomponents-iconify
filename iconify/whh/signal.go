package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Signal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832.5 1024q-26.5 0-45.5-18.5T768 960V64q0-26 19-45t45.5-19t45 18.5T896 64v896q0 27-18.5 45.5t-45 18.5zm-256 0q-26.5 0-45.5-18.5T512 960V320q0-27 19-45.5t45.5-18.5t45 19t18.5 45v640q0 27-18.5 45.5t-45 18.5zm-256 0q-26.5 0-45.5-18.5T256 960V576q0-27 19-45.5t45.5-18.5t45 18.5T384 576v384q0 27-18.5 45.5t-45 18.5zM64 1024q-27 0-45.5-18.5T0 960V832q0-27 18.5-45.5T64 768t45.5 18.5T128 832v128q0 27-18.5 45.5T64 1024z"/>`),
		g.Group(children),
	)
}