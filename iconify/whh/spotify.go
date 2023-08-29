package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spotify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M775 951q-52-84-138.5-133.5T448 768q-81 0-154 32.5T168 890Q89 818 44.5 720T0 512q0-104 40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199q0 138-67.5 254T775 951zM316 658q67-18 132-18q120 0 229 54l28-57q-122-61-257-61q-73 0-148 20zm-49-186q92-24 181-24q165 0 315 75l28-57q-163-82-343-82q-98 0-198 27zm181-280q-126 0-247 32l16 63q117-31 231-31q210 0 400 95l29-58Q675 192 448 192zm208 811q-72 21-144 21q-126 0-238-59q74-69 174-69q63 0 117.5 29t90.5 78z"/>`),
		g.Group(children),
	)
}