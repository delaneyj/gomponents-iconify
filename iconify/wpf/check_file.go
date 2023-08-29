package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M4 0C1.8 0 0 1.8 0 4v18c0 2.2 1.8 4 4 4h11.906l-2-2H4c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h7.313c.7.2.687 1.1.687 2v3c0 .6.4 1 1 1h3c1 0 2 0 2 1v9.813l.188.187L20 16.312V8c0-1.1-.988-2.112-2.688-3.813c-.3-.2-.512-.487-.812-.687c-.2-.3-.488-.513-.688-.813C14.113.988 13.1 0 12 0H4zm20.125 14a.777.777 0 0 0-.438.313l-5.28 7.78l-2.407-2.5c-.3-.3-.706-.3-.906 0l-.906.907c-.3.3-.3.706 0 .906l3.718 3.688c.2.2.482.406.782.406c.3 0 .612-.2.812-.5l6.406-9.406c.2-.1.082-.582-.218-.781l-1.094-.72a.58.58 0 0 0-.469-.093z"/>`),
		g.Group(children),
	)
}