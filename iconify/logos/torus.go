package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Torus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<defs><linearGradient id="svgIDa" x1="-.012%" x2="100.007%" y1="50.013%" y2="50.013%"><stop offset="0%" stop-color="#A34CB4"/><stop offset="9.192%" stop-color="#954DB2"/><stop offset="71.69%" stop-color="#3C52A1"/><stop offset="100%" stop-color="#19549B"/></linearGradient><linearGradient id="svgIDb" x1="-3.884%" x2="83.237%" y1="172.005%" y2="-50.934%"><stop offset="0%" stop-color="#FF0264"/><stop offset="49.43%" stop-color="#FE744A"/><stop offset="81.9%" stop-color="#FDBC39"/><stop offset="100%" stop-color="#FDDF31"/></linearGradient></defs><path fill="url(#svgIDa)" d="M128 0C81.6 0 41 12.5 18.5 31.3c.1.5-.4 1.5-.2 2C20.9 45.7 32 57 48.8 65.4c.4.2 1.5-.4 2-.2c1.9-1.3 25.6-16.9 77.2-16.9c51.3 0 75.2 15.4 77.3 16.9c.4-.2 1.4.2 1.8.1c17.1-8.5 28.4-19.8 30.8-32.2c.1-.4-.5-1.2-.4-1.6C215 12.5 174.4 0 128 0"/><path fill="url(#svgIDb)" d="M237.5 31.3c-4.3 27.6-51.7 49.3-109.5 49.3S22.8 58.9 18.5 31.3C6.8 41.1 0 52.6 0 64.9c0 35.8 57.3 64.9 128 64.9s128-29.1 128-64.9c0-12.3-6.8-23.8-18.5-33.6"/>`),
		g.Group(children),
	)
}