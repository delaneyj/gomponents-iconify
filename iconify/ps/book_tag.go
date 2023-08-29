package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookTag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M88 512h341v-43H88q-17 0-30-12.5T45 427q0-18 13-30.5T88 384h341V0H67Q39 0 21 18.5T3 64v363q0 35 25 60t60 25zM45 64q0-21 22-21h320v298H88q-19 0-43 11V64zm22 363q0 21 21 21h341v-43H88q-21 0-21 22zM344 85H109v128h235V85zm-43 86H152v-43h149v43z"/>`),
		g.Group(children),
	)
}