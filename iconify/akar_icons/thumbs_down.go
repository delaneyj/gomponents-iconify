package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12.395 18.218l-.23 2.369c-.091.952-.98 1.598-1.878 1.366c-1.351-.35-2.3-1.605-2.3-3.044v-3.035c0-.675 0-1.013-.146-1.26a1.018 1.018 0 0 0-.333-.345c-.24-.151-.567-.151-1.22-.151h-.396c-1.703 0-2.554 0-3.078-.39a2.073 2.073 0 0 1-.78-1.208c-.146-.65.181-1.463.836-3.087l.327-.81c.188-.468.265-.975.225-1.48c-.232-2.874 2.047-5.295 4.833-5.135l10.424.598c1.139.065 1.708.098 2.222.553c.515.455.612.924.805 1.861a14.317 14.317 0 0 1-.055 6.037c-.283 1.248-1.475 1.92-2.706 1.76c-3.264-.42-6.223 2.019-6.55 5.4Z"/><path d="m17 12.5l.137-.457c.887-2.956.84-6.115-.137-9.043"/></g>`),
		g.Group(children),
	)
}