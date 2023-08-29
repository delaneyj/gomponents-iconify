package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radiation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.912 8.531L7.121 3.877a.501.501 0 0 0-.704-.166a9.982 9.982 0 0 0-4.396 7.604a.505.505 0 0 0 .497.528l5.421.09a4.042 4.042 0 0 1 1.973-3.402zm8.109-4.51a.504.504 0 0 0-.729.151L14.499 8.83a4.03 4.03 0 0 1 1.546 3.112l5.419-.09a.507.507 0 0 0 .499-.53a9.986 9.986 0 0 0-3.942-7.301zm-4.067 11.511a4.015 4.015 0 0 1-1.962.526a4.016 4.016 0 0 1-1.963-.526l-2.642 4.755a.5.5 0 0 0 .207.692A9.948 9.948 0 0 0 11.992 22a9.94 9.94 0 0 0 4.396-1.021a.5.5 0 0 0 .207-.692l-2.641-4.755z"/><circle cx="12" cy="12" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}