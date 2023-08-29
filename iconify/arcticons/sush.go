package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.015 27.555s-1.623 2.03.288 3.317c-1.287.602-1.807 3.552.727 3.843c-2.326.727-1.184 4.86 2.41 3.49c-.27 1.35.935 3.343 3.282 1.287c1.973 1.37 3.697 1.766 7 .208c1.017 1.765 3.344 1.724 4.507 0c2.763 1.848 5.878 1.454 8.516-.167c3.033 1.974 4.944-1.454 3.905-2.783c3.033-.498 3.158-3.863 1.08-4.258c1.247-2.305 1.101-4.923-1.163-6.169c2.867-1.018 2.119-2.866-.353-3.178c1.08-1.496-1.1-3.157-1.1-3.157"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.879 33.282c1.252-.52 4.402-4.262 3.442-5.775c-.686-1.08-2.046-.02-4.43 0M7.436 28.884c-2.3 1.858-3.95 1.842-3.936 3.089c.017 1.553 2.593 2.05 4.171 2.108m30.626-20.366c4.383 2.243 1.08 6.148-1.184 6.273c-3.516.193-7.581-1.641-13.771-1.641c-4.902 0-11.3 2.306-12.608 8.142c-1.87-2.866.104-7.789.104-7.789c-2.168 2.887-2.247 7.962-1.704 9.638c-6.397-3.697-2.866-12.9 3.158-12.9c4.88 0 7.604 3.27 7.604 3.27"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.888 10.848s-6.879-2.486-12.754-1.37c-5.359 1.017-7.889 2.87-8.14 7.791"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.42 15.475c2.94.975 7.15 1.652 7.57-.872c.421-2.524-2.398-2.976-2.398-2.976c2.804-1.48-1.698-3.115-.374-4.595c-2.586-.234-3.319.561-3.35 2.158c-1.885-1.052-3.432-.554-2.98 1.659c-2.695-.094-2.573.823-2.41 2.025c.183 1.353 2.311 2.061 3.941 2.601Z"/><circle cx="17.972" cy="30.498" r="2.342" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.343" cy="30.498" r="2.342" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.65 32.394h-2.985a.662.662 0 0 0-.662.662c0 1.036.84 1.877 1.877 1.877h.555c1.037 0 1.878-.84 1.878-1.877a.662.662 0 0 0-.663-.662Z"/>`),
		g.Group(children),
	)
}