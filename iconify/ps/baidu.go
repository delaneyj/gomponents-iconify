package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Baidu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M109 174q2 29-12 51t-36 23t-38.5-18.5T4 181t11.5-51.5T51 106q22-2 39 18t19 50zM308 11q-21-5-41.5 11.5T239 67q-7 29 3 53t31 29t41-11.5T341 93q6-22-1.5-42t-17-29T308 11zm66 128q-23 1-38 21t-15 52q1 31 16 47t39 16q53-2 52-65q0-26-12.5-43.5t-24-22.5t-17.5-5zM160 2q-22 0-36.5 20.5T109 72t14.5 49.5T160 142q21 0 35.5-20.5T210 72t-14.5-49.5T160 2zm-7 218q-25 37-70 76q-51 42-51 81q0 29 16.5 55T98 458q26 0 60-5.5t52-5.5t53.5 7.5T325 462q33 0 51.5-24.5T395 385q0-23-9.5-39.5T347 303q-47-39-81-87q-18-26-56-26q-36 0-57 30z"/>`),
		g.Group(children),
	)
}