package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wulkanowy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.493 42.5h-6.748l-4.343-5.775l-1.841-6.183l-5.554-2.81l-.934-4.112l-1.689-2.458l1.406-2.73l2.202-.056l2.078-1.247l5.218 10.143l.398 3.71l9.81 10.888Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.308 42.5h-9.674l-2.69-6.411l2.186-5.236l-1.603-3.285l2.408-4.63l.597.186l2.474 7.105l4.495 1.53l2.177 6.344Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.261 40.555L14.78 42.5H5.504l5.8-9.075l2.593-11.281l3.47-2.223a9.414 9.414 0 0 1 .184-2.258c.282-.459 1.588-.279 1.588-.279a.84.84 0 0 1 .415.922l-3.555 9.143l.841 4.336l-2.072 4.495Zm4.741-24.295c.453-1.034.104-1.354.104-1.354c-2.192-.451-3.797-3.55-.742-5.21c-.39-.922.787-1.648 1.542-1.55c.482-.732.879-.853 1.593-.61c1.458-2.014 3.616-1.92 6.494-1.34c2.496-1.735 4.19.168 3.74 1.465c2.364.398 1.593 3.503-.253 3.083c.006.802-.323 1.243-1.188 1.108c-.625 1.968-2.68 2.569-5.963 1.994a.975.975 0 0 1-1.27.697c-.037.944-.127 1.859-2.508 1.581a.974.974 0 0 1-.903.788c-.498-.012-.833-.225-.646-.652Z"/>`),
		g.Group(children),
	)
}