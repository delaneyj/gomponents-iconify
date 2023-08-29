package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Github(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M510 383q23 0 39 21t15 53t-15 52t-39 22t-38-22t-16-52t16-53t38-21zm186-193q29 31 46 69t16 90q0 74-18 125t-48 84t-64 52t-70 28t-64 10l-44 2H308q-15 0-44-2t-63-10t-70-28t-65-52t-47-84T0 349q0-51 17-90t45-69q-3-6-3-24t1-42t9-56T88 8q22 3 51 14q25 9 59 25t77 46q18-5 46-8t58-2l58 2q28 1 46 8q42-29 77-46t59-25q29-11 51-14q12 30 19 60t9 56t2 42t-4 24zM380 614q58 0 109-6t88-23t59-51t21-90q0-27-10-52t-33-45q-19-18-44-24t-56-5t-64 4t-70 3h-2q-36 0-70-3t-64-4t-55 5t-44 24q-23 20-33 45t-11 52q0 57 22 90t58 51t88 23t109 6h2zM248 383q23 0 39 21t15 53t-15 52t-39 22t-38-22t-16-52t16-53t38-21z"/>`),
		g.Group(children),
	)
}