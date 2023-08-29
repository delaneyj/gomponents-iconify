package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alitelecom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.112 25.656a13.032 13.032 0 0 1 6.444-11.228a12.693 12.693 0 0 1 12.888 0a13.032 13.032 0 0 1 6.444 11.228"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.183 25.656a17.845 17.845 0 0 1 8.104-15.133a15.184 15.184 0 0 1 16.207 0a17.845 17.845 0 0 1 8.104 15.133"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.95 22.239a2.21 2.21 0 0 1-2.246 2.245h0a2.21 2.21 0 0 1-2.245-2.245h0a2.21 2.21 0 0 1 2.245-2.246h0a2.271 2.271 0 0 1 2.246 2.246Zm9.568 0a2.21 2.21 0 0 1-2.246 2.245h0a2.21 2.21 0 0 1-2.245-2.245h0a2.21 2.21 0 0 1 2.245-2.246h0a2.21 2.21 0 0 1 2.246 2.246ZM24 31.052c2.043-2.042 8.434 3.342 14.169 3.342c3.603 0 4.967-4.394 4.967-4.394s3.327 9.82-10.14 9.82A9.944 9.944 0 0 1 24 34.785a9.944 9.944 0 0 1-8.995 5.036C1.537 39.82 4.864 30 4.864 30s1.364 4.394 4.967 4.394c5.736 0 12.126-5.384 14.169-3.341Z"/>`),
		g.Group(children),
	)
}