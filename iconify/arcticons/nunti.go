package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nunti(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.868 26.587a14.873 14.873 0 0 1-29.738-.093M24 15.451s1.329 1.123-17.052 12.364L4.5 26.17s.098-2.701.52-6.3s9.976-9.395 12.932-11.332S21.875 6.961 24 6.961"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 15.451s-1.329 1.123 17.052 12.364L43.5 26.17s-.098-2.701-.52-6.3s-9.976-9.395-12.932-11.332S26.125 6.961 24 6.961"/><ellipse cx="24" cy="31.627" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.646" ry="5.379"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.904 23.834v4.36m14.192-4.36v4.36M21.7 14.338c.09-.073-.033-7.367-.033-7.367m4.697 7.423c.091-.074-.031-7.423-.031-7.423M21.7 14.338s1.18.593-17.2 11.833m21.864-11.777S25.12 14.93 43.5 26.17"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.7 14.338a3.227 3.227 0 0 1 4.664.056m-7.01 17.259a8.68 8.68 0 0 0 9.292 0"/>`),
		g.Group(children),
	)
}