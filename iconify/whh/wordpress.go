package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wordpress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 192q0 14 1 24.5t4.5 21t6 17t10 17t10.5 15t14.5 18.5t16.5 19q22 28 28.5 45.5T861 410q-7 34-16 60l-77 202l-83-188q-9-22-37-117.5T620 264q0-14 10-21q22-18 42-19v-32H384v32q9 1 14 6t9.5 11.5t7.5 9.5q14 12 33 58l32 107l-64 256l-132-349q-20-51-20-62t11-19q24-18 45-18v-32H113q71-90 175.5-141T512 0q95 0 182 33.5T850 128q-39 0-60.5 16T768 192zM66 261q25 29 60 123l194 512h64l128-384l160 384h64l151-390q6-17 24-53.5t30.5-70T957 322q3-40 3-58q64 116 64 248q0 139-68.5 257T769 955.5T512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512q0-134 66-251z"/>`),
		g.Group(children),
	)
}