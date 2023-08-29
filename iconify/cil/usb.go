package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M320 224h32v40h-80V78.627l17.941 17.941l22.627-22.627L256 17.373l-56.568 56.568l22.627 22.627L240 78.627V264h-72v-42.341a56 56 0 1 0-32 0V296h104v90.341a56 56 0 1 0 32 0V296h112v-72h40V120H320Zm-192-56a24 24 0 1 1 24 24a24.027 24.027 0 0 1-24-24Zm152 272a24 24 0 1 1-24-24a24.027 24.027 0 0 1 24 24Zm72-288h40v40h-40Z"/>`),
		g.Group(children),
	)
}