package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.392 18.372c-.574-1.234-.99-2.013.77-3.558c.85-10.428 14.722-14.074 18.849-.48a4.515 4.515 0 0 0 .131 1.77C37.216 21.893 38 30.26 37.57 34.816c1.064 2.146 2.238 3.142 3.558 2.597c.728 1.317 2.602 4.23-.096 4.712c-9.003.873-8.526.023-9.424-.577c-1.572 1.425-13.656.577-20.484.865c-1.364.005-3.066-.803-2.885-2.308c-2.748-1.52-4.286-10.37.9-9.294c.374-4.734 1.325-7.83 5.255-12.44Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.01 42.373c3.362-2.703 3.455-5.657 3.628-8.282M8.237 40.106c.63-4.43 4.859-1.138 6.873-6.15"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.137 30.812a10.155 10.155 0 0 0 1.927 6.34m20.542 4.397a11.23 11.23 0 0 0 5.963-6.732M22.08 12.638a.76.76 0 0 1-1.142-.105a.994.994 0 0 1 .379-1.223a.761.761 0 0 1 1.143.103a.993.993 0 0 1-.378 1.223m6.948.082a.911.911 0 0 1-1.13-.35a.708.708 0 0 1 .118-1.06a.91.91 0 0 1 1.13.35a.708.708 0 0 1-.117 1.059m-3.868 3.179l-.723-1.142l-.627-1.198l1.35-.055l1.351.055l-.627 1.198Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.163 15.847l-.024 1.106l-1.395.89m1.395-.889l1.594.883m-3.878-8.794l1.033 1.13m2.746.005l1.01-.89"/>`),
		g.Group(children),
	)
}