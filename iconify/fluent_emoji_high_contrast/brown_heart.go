package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrownHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<mask id="fluentEmojiHighContrastBrownHeart0" width="28" height="25" x="2" y="4" maskUnits="userSpaceOnUse" style="mask-type:alpha"><path fill="#C4C4C4" d="M6 5.08c4.665-2.333 8.5.5 10 2.5c1.5-2 5.335-4.833 10-2.5c6 3 4.5 10.5 0 15c-2.196 2.195-6.063 6.062-8.891 8.213a1.764 1.764 0 0 1-2.186-.04C12.33 26.16 8.165 22.243 6 20.078c-4.5-4.5-6-12 0-15Z"/></mask><g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="2" d="M15.2 8.18a1 1 0 0 0 1.6 0a7.973 7.973 0 0 1 3.373-2.542c1.492-.575 3.322-.693 5.38.336c2.593 1.296 3.551 3.511 3.387 5.976c-.169 2.53-1.536 5.311-3.647 7.422c-2.213 2.214-6.027 6.024-8.79 8.125a.764.764 0 0 1-.952-.023c-2.553-2.06-6.69-5.947-8.844-8.102c-2.111-2.11-3.478-4.893-3.647-7.422c-.164-2.465.794-4.68 3.387-5.976c2.058-1.03 3.888-.911 5.38-.336A7.973 7.973 0 0 1 15.2 8.179Z"/><g stroke-linecap="round" stroke-width=".5" mask="url(#fluentEmojiHighContrastBrownHeart0)"><path d="m2.043-19.454l28.62 28.62M2.043-4.454l28.62 28.62m-28.62-13.62l28.62 28.62m-28.62-58.62l28.62 28.62M2.043-4.454l28.62 28.62m-28.62-13.62l28.62 28.62m-28.62-53.62l28.62 28.62M2.043.546l28.62 28.62m-28.62-13.62l28.62 28.62m-28.62-23.62l28.62 28.62M2.043-9.454l28.62 28.62M2.043 5.546l28.62 28.62m-28.62-8.62l28.62 28.62"/></g></g>`),
		g.Group(children),
	)
}