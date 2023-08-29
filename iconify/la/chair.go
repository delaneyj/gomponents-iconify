package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11 4v12.438c-.102.027-.223.066-.313.093c-.484.149-.82.305-1.062.469a1.986 1.986 0 0 0-.313.25c-.078.082-.187.25-.187.25L9 17.719V20h1v8h2v-8h8v8h2v-8h1v-2.281l-.125-.219s-.11-.168-.188-.25a1.986 1.986 0 0 0-.312-.25c-.242-.164-.578-.32-1.063-.469c-.09-.027-.21-.066-.312-.093V4h-2v1h-6V4zm2 3h2v9.031c-.758.02-1.438.04-2 .094zm4 0h2v9.125c-.563-.055-1.242-.074-2-.094z"/>`),
		g.Group(children),
	)
}