package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimeFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0zm0 469.3c-117.8 0-213.3-95.5-213.3-213.3c0-117.8 95.5-213.3 213.3-213.3c117.8 0 213.3 95.5 213.3 213.3c0 117.8-95.5 213.3-213.3 213.3zm-21.3-256L181.3 160l-32 32l128 128V85.3h-42.7v128z"/>`),
		g.Group(children),
	)
}