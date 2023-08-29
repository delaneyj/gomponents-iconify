package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RnzEpaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.513 22.73v-9.657h3.138c1.81 0 3.26 1.449 3.26 3.26s-1.45 3.259-3.26 3.259h-3.138m3.376-.002l3.021 3.14m3.029 0v-9.657l6.397 9.657v-9.657m3.029 0h6.397l-6.397 9.657h6.397M8.5 32.645h3.38M8.5 25.886h3.38M8.5 29.265h2.197M8.5 25.886v6.759m8.188 0v-6.76h2.197c1.267 0 2.28 1.015 2.28 2.282s-1.013 2.281-2.28 2.281h-2.197m21.122-.591c0-.93.76-1.69 1.69-1.69h0m-1.69 0v4.478M35.9 31.8c-.254.507-.846.845-1.437.845h0c-.93 0-1.69-.76-1.69-1.69v-1.098c0-.93.76-1.69 1.69-1.69h0c.93 0 1.69.76 1.69 1.69v.591h-3.38m-6.86.508c0 .929-.76 1.69-1.69 1.69h0c-.93 0-1.69-.761-1.69-1.69v-1.099c0-.93.76-1.69 1.69-1.69h0c.93 0 1.69.76 1.69 1.69m0 2.788v-4.478m-13.222 2.174h2.399m12.595.615c0 .929.76 1.69 1.69 1.69h0c.93 0 1.69-.761 1.69-1.69v-1.099c0-.93-.76-1.69-1.69-1.69h0c-.93 0-1.69.76-1.69 1.69m0-1.69v6.76"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/>`),
		g.Group(children),
	)
}