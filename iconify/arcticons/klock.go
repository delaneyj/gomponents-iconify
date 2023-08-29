package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Klock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.748 33.987c-2.34-2.075-2.035-5.9-2.035-5.9a1.869 1.869 0 0 0 0-3.738a5.162 5.162 0 0 0-3.25-9.14v-.002h-.04v.003a5.142 5.142 0 0 0-2.271.545m-1.612 1.26a5.162 5.162 0 0 0 .633 7.334a1.869 1.869 0 0 0 0 3.737s.428 5.405-3.663 6.851v2.794H9a1.385 1.385 0 0 0-1.39 1.389v2.991A1.385 1.385 0 0 0 9 43.5h12.832"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.097 15.71a2.205 2.205 0 1 0 0 4.41s.507 11.571-4.325 13.277v3.297h-.602a1.635 1.635 0 0 0-1.64 1.639v3.529a1.635 1.635 0 0 0 1.64 1.638h17.58a1.635 1.635 0 0 0 1.64-1.639v-3.528a1.635 1.635 0 0 0-1.64-1.639h-.605v-3.297a3.83 3.83 0 0 1-1.337-.824m-1.447-2.12c-1.82-3.985-1.54-10.333-1.54-10.333a2.205 2.205 0 1 0 0-4.41"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.097 15.71s-3.638.689-3.851-2.425s-.653-8.785-.653-8.785m12.228 11.21s3.638.689 3.852-2.425s.652-8.785.652-8.785"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.593 4.5h2.956l1.211 4.52h1.509l1.2-4.478h1.302"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.325 4.5H35.37l-1.211 4.52H32.65l-1.2-4.478h-1.68"/>`),
		g.Group(children),
	)
}