package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Woman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M192 64q0 26-19 45t-45 19t-45-19t-19-45t19-45t45-19t45 19t19 45zM64 256v235q0 9 6 15t15 6q10 0 16-6t6-15v-43h42v43q0 9 6 15t16 6q9 0 15-6t6-15V192q0-17-12.5-30T149 149h-42q-18 0-30.5 13T64 192v64zM43 149q-18 0-30.5 13T0 192v149q17 0 30-12.5T43 299V149zm213 43q0-17-12.5-30T213 149v150q0 17 13 29.5t30 12.5V192z"/>`),
		g.Group(children),
	)
}