package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Remini(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.262 20.59v-5.112h4.257M15.32 29.994v1.533l2.07-.028m2.015-1.505l-.199 1.705l11.607 1.562l.34-2.557M20.598 20.59l.85-5.851l11.892 1.79l-.625 4.062M8 28.548v-8.162h2.65c1.529 0 2.752 1.224 2.752 2.754s-1.223 2.755-2.752 2.755H8m2.851-.009l2.558 2.659m7.399-3.261c0-1.122.917-2.04 2.038-2.04s2.039.918 2.039 2.04v3.265m-4.077-5.305v5.305m4.077-3.265c0-1.122.917-2.04 2.038-2.04s2.039.918 2.039 2.04v3.265"/><circle cx="30.845" cy="20.693" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.845 23.142v5.407m-12.22-1.022c-.305.612-1.019 1.02-1.732 1.02a2.046 2.046 0 0 1-2.039-2.04V25.18c0-1.122.917-2.04 2.039-2.04s2.038.918 2.038 2.04v.714h-4.077m22.169 2.656v-3.368c0-1.122-.918-2.04-2.039-2.04s-2.038.918-2.038 2.04v3.367m0-3.367v-2.04"/><circle cx="39.185" cy="20.693" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.185 23.142v5.407"/>`),
		g.Group(children),
	)
}