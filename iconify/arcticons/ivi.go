package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ivi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m24.11 27.48l-3.36-5.96l-3.34 5.96h6.7zm13.58 6.85l5.81-3.43l-5.81-3.45v6.88zm0-6.89l5.81 3.45v-6.88l-5.81 3.43zm-10.19 6l3.34-5.96h-6.7l3.36 5.96z"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="m27.48 21.52l-3.34 5.96h6.7l-3.36-5.96zm3.36 5.93l3.37-5.96h-6.73l3.36 5.96zm6.85-6.89l5.81-3.42v-.03l-5.81-3.42v6.87zm0 6.88v.02l5.81-3.45l-5.81-3.45v6.88zM4.5 20.55l5.81-3.42v-.03L4.5 13.68v6.87zm0 6.87v.03L10.31 24L4.5 20.55v6.87zm16.26 6.02h6.73l-3.36-5.96l-3.37 5.96z"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="m17.41 27.48l3.37 5.96l3.33-5.96h-6.7zm-3.35-5.99l3.37 5.96l3.33-5.96h-6.7zM4.5 34.32l5.81-3.42l-5.81-3.45v6.87zm5.81-3.43v-6.88L4.5 27.44l5.81 3.45z"/>`),
		g.Group(children),
	)
}