package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shelly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.5 25.799v2.696a2.003 2.003 0 0 1-1.997 1.998h0a1.676 1.676 0 0 1-1.399-.6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.5 22.503V25.8a2.003 2.003 0 0 1-1.997 1.997h0a2.003 2.003 0 0 1-1.998-1.997v-3.296m-7.047-2.997v6.991a.944.944 0 0 0 .999.999h.3m2.225-7.99v6.991a.944.944 0 0 0 .998.999h.3m-7.347-.999a1.933 1.933 0 0 1-1.697.999h0a2.003 2.003 0 0 1-1.998-1.998v-1.298a2.003 2.003 0 0 1 1.998-1.997h0a2.003 2.003 0 0 1 1.997 1.997v.699h-3.995m-6.22-5.393v7.99m0-3.296a2.003 2.003 0 0 1 1.998-1.997h0a2.003 2.003 0 0 1 1.997 1.997v3.296M8.6 26.598a2.381 2.381 0 0 0 1.997.899h1.199a2.003 2.003 0 0 0 1.997-1.998h0a2.003 2.003 0 0 0-1.997-1.997h-1.299A2.003 2.003 0 0 1 8.5 21.505h0a2.003 2.003 0 0 1 1.997-1.998h1.199a2.144 2.144 0 0 1 1.997.899"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}