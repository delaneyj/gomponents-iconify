package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bleentoro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.37 13.08h6.24V5.51"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.17 10.38h2.46a.76.76 0 0 1 .74.78v4.93a.76.76 0 0 1-.74.78h-2.46a.76.76 0 0 1-.74-.78v-4.93a.76.76 0 0 1 .74-.78Zm2.28 23.24h4.83a.76.76 0 0 1 .74.78v2.4a.76.76 0 0 1-.74.78h-4.83a.76.76 0 0 1-.74-.78v-2.4a.76.76 0 0 1 .74-.78Zm20.9-25.55h2.41a.76.76 0 0 1 .74.78v2.37a.76.76 0 0 1-.74.78h-2.41a.76.76 0 0 1-.74-.78V8.85a.76.76 0 0 1 .74-.78Zm-6.91 7.06h2.41a.76.76 0 0 1 .74.78v2.36a.76.76 0 0 1-.74.78h-2.41a.76.76 0 0 1-.74-.78v-2.36a.76.76 0 0 1 .74-.78ZM32 36h2.41a.76.76 0 0 1 .74.78v2.37a.75.75 0 0 1-.74.77H32a.75.75 0 0 1-.74-.77V36.8A.76.76 0 0 1 32 36ZM11.84 10.38V5.51m.86 30.16H5.51m8.67 1.92v4.91M34.24 40v2.5M37.5 9.35h5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.61 10.41h-5.35v4.72M23.58 20v-1.77h-5.17v-5.14m9.69 5.97v1M17.29 30.92v2.7M23.57 31v6.9h7.7m2.88-7.01V36m0-3.29h8.35M35.12 12v5h2.61v3m-12.86 7a2 2 0 0 1-1.74 1h0a2 2 0 0 1-2-2v-1.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2v.65h-4M30.87 27a2 2 0 0 1-1.74 1h0a2 2 0 0 1-2-2v-1.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2v.65h-4M18.32 20v7a1 1 0 0 0 1 1h.3m17.48 0v-3.3a2 2 0 0 0-2-2h0a2 2 0 0 0-2 2V28m0-3.3v-2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".87" d="M14.2 24a2 2 0 1 1 0 4h-3.3v-8h3.3a2 2 0 1 1 0 4Zm0 0h-3.3"/>`),
		g.Group(children),
	)
}