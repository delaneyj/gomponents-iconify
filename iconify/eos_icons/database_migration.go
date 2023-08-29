package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseMigration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 11.021c0 1.013-2.462 1.833-5.5 1.833s-5.5-.82-5.5-1.833V9.188c0 1.012 2.462 1.833 5.5 1.833s5.5-.82 5.5-1.833"/><path fill="currentColor" d="M18.5 6.137c-3.035 0-5.5-.825-5.5-1.833v3.64c0 1.009 2.465 1.834 5.5 1.834S24 8.953 24 7.944v-3.64c0 1.008-2.465 1.833-5.5 1.833Z"/><path fill="currentColor" d="M24 1.833C24 .821 21.538 0 18.5 0S13 .82 13 1.833v1.223c0 1.012 2.462 1.833 5.5 1.833s5.5-.82 5.5-1.833M11 22.021c0 1.013-2.462 1.833-5.5 1.833S0 23.034 0 22.021v-1.833C0 21.2 2.462 22.02 5.5 22.02s5.5-.82 5.5-1.833Z"/><path fill="currentColor" d="M5.5 17.137c-3.035 0-5.5-.825-5.5-1.833v3.64c0 1.009 2.465 1.834 5.5 1.834s5.5-.825 5.5-1.834v-3.64c0 1.008-2.465 1.833-5.5 1.833Z"/><path fill="currentColor" d="M11 12.833C11 11.821 8.538 11 5.5 11S0 11.82 0 12.833v1.223c0 1.012 2.462 1.833 5.5 1.833s5.5-.82 5.5-1.833ZM20 21h-2v-2.59L15.41 21L14 19.59L16.59 17H14v-2h6v6z"/>`),
		g.Group(children),
	)
}