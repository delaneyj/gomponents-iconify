package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForAnguilla(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#2a5f9e" d="M12 12h20.4v20.4H12z"/><path fill="#fff" d="M18.9 8.1V14h-7.3L26 32h6v-7.5z"/><path fill="#2a5f9e" d="M32 2v30H2c0 16.6 13.4 30 30 30s30-13.4 30-30S48.6 2 32 2z"/><path fill="#fff" d="M11 18.9H5c-1.9 4-3 8.4-3 13.1h12V13.9h-3v5"/><path fill="#ed4c5c" d="M32 27.1L19 11h-6l17 21h2z"/><path fill="#fff" d="M18.9 5v6H11v3h21V2c-4.7 0-9.1 1.1-13.1 3z"/><path fill="#ed4c5c" d="M32 5H18.9c-6 2.9-11 7.9-13.9 13.9V32h6V11h21V5z"/><path fill="#fff" d="M50 33.4c-1.6 0-3-1.4-3-1.4s-1.4 1.4-3 1.4c-1.5 0-4-1.4-4-1.4s.4 8.8 2 11.8c1.8 3.2 5 6.2 5 6.2s3.2-3 5-6.2c1.6-3 2-11.8 2-11.8s-2.5 1.4-4 1.4"/><path fill="#b4d7ee" d="M42.9 43.3c1.2 2.1 3.1 4.3 4.1 5.3c1-1 2.9-3.2 4.1-5.3h-8.2"/><g fill="#ff8736"><path d="M47 35.5c-.9 1-.4 3.6-1.3 2.7c-.9-1-.9-2.5 0-3.5s2.4-1 3.3 0c.9 1-1.1-.1-2 .8M44.8 40c1.3.3 3.2-1.5 2.9-.2c-.3 1.3-1.6 2.1-2.9 1.7c-1.3-.3-2-1.7-1.7-3c.3-1.3.4 1.1 1.7 1.5"/><path d="M49.6 39.7c-.3-1.3-2.8-2.1-1.6-2.5c1.3-.3 2.6.4 2.9 1.7c.3 1.3-.4 2.6-1.7 3c-1.2.4.7-.9.4-2.2"/></g>`),
		g.Group(children),
	)
}