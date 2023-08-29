package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Motorway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#699635" d="M0 20.4s1.9-3.7 4.6-5.5c6.4-4.2 4.8 9.1 14.5 3.3c3.8-2.2 4.7 2.9 6.9 3.4c4.7 1.2 12.3.6 12.3.6s11.2-2.4 14.8-5.2c3-2.4 5 .2 8.1-1.1c1.3-.6 2.8 0 2.8 0v15.8H0V20.4"/><path fill="#83bf4f" d="M27 21s-6.8.3-12.1-3.7c-3.7-2.8-5.5 5.4-8.8 5.4c-3.4 0-6-4.6-6-4.6v35.2L27 21m10 0s2.5.5 9.4-2.6c3.2-1.5 4.1 5.7 5.9 5.7c1.6 0 4.2-3.8 6.6-3.9c3.2-.2 5 1.6 5 1.6v31.4L37 21z"/><path fill="#ffdd7d" d="M38.1 21H25.9L0 46.1V64h64V46.1z"/><path fill="#666" d="M37 21H27L0 52.2V64h64V52.2z"/><path fill="#fff" d="M64 58v-3.7L36.6 21h-.4zM0 58v-3.7L27.4 21h.4zm33.8 6h-3.6l.4-11.9h2.8zm-.7-18.7h-2.2l.3-10h1.6zm-.4-12.5h-1.4l.1-4.3h1.2zm-.2-5.5h-1l.1-2.5h.8zm-.1-3.2h-.8l.1-1.5h.6zm-.1-2h-.6V21h.6z"/>`),
		g.Group(children),
	)
}