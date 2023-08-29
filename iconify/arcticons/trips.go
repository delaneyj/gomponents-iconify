package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trips(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5c-5.936 0-10.748 4.84-10.748 10.81c0 8.465 8.21 16.313 10.354 18.807a.654.654 0 0 0 .993 0c2.094-2.477 10.149-10.342 10.149-18.807C34.748 9.34 29.936 4.5 24 4.5Zm0 14.865c-3.587.007-5.389-4.352-2.856-6.905c2.533-2.553 6.871-.75 6.871 2.858c0 2.23-1.798 4.039-4.015 4.039v.008ZM18.88 43.5v-7.638h1.241a3.83 3.83 0 0 1 3.82 3.82h0a3.83 3.83 0 0 1-3.82 3.818h-1.24Zm-1.765 0l-2.482-7.638l-2.578 7.638m.859-2.578h3.342m19.689 0h0a2.567 2.567 0 0 1-2.578 2.578h0a2.567 2.567 0 0 1-2.577-2.578v-2.577a2.567 2.567 0 0 1 2.578-2.578h0c1.431 0 2.482 1.146 2.482 2.578h0M29.816 43.5l-2.482-7.638l-2.578 7.638m.859-2.578h3.341"/>`),
		g.Group(children),
	)
}