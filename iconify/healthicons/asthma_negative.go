package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsthmaNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsAsthmaNegative0)"><path d="M15.5 23a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm23.742-3.38l.463-1.72l-2.536-.676l-.349 1.294l2.422 1.102Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM5.918 6.044C6.275 6.015 6.636 6 7 6c7.18 0 13 5.82 13 13v1.697l2.796 4.194C23.682 26.22 22.729 28 21.13 28H20v2a1 1 0 0 1-1 1h-1a2 2 0 0 0-2 2h3a1 1 0 0 1 1 1v3a4 4 0 0 1-1.917 3.416c-.715.436-1.484.494-2.16.405c-.666-.088-1.303-.324-1.826-.549L10 38.517V43a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V7.041a1 1 0 0 1 .918-.997Zm30.802 8.99a1 1 0 0 0-1.223.706l-.493 1.83a1 1 0 0 0-1.372.65L31.247 27H28a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h6.204a5 5 0 0 0 4.82-3.673l2.94-10.678a1 1 0 0 0-.551-1.177l-.294-.132l.777-2.888a1 1 0 0 0-.708-1.226l-4.468-1.192ZM26 35v-7h-2v7h2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsAsthmaNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}