package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlmostAHero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.62 16.284C35.104 5.029 23.789 6.235 21.659 9.368c-2.921 4.296-1.507 11.442 2.105 13.016c1.208.526 2.856-.881 6.529.666c.859 1.668 4.836 3.39 5.928.837c-1.747.286-4.835-1.49-5.08-2.892"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.829 8.317C20.47 6.103 15.2 5.989 11.335 7.564a4.87 4.87 0 0 1 2.205 1.088s-5.527 2.929-5.985 6.476a11.15 11.15 0 0 1 3.98.712m7.245 21.278c2.721 3.465 8.764 4.381 8.764 4.381a8.093 8.093 0 0 0 2.978-2.806c5.498 2.434 8.334 1.976 11.799 1.174c.057-4.926-1.06-6.93-1.06-6.93a5.17 5.17 0 0 0 2.147 0c.917-5.069-5.297-12.142-8.877-13.202c-.286 1.49-.918 2.459-2.378 2.676"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.277 27.639c1.632 1.002 8.305 1.117 9.278.802c-1.002-1.661-1.832-4.21-4.639-4.21s-4.64 3.408-4.64 3.408Zm1.833-10.711c1.231-.945 5.327-.916 4.926 3.265c-2.091.515-5.012-.029-4.926-3.265Zm5.87 3.036c.115-2.12 1.49-4.353 4.897-3.35a3.442 3.442 0 0 1-4.897 3.35Zm-18.586-1.232l10.625 11.083l-1.947 2.377l-10.997-11.025l-5.241 4.925m5.779-4.386l2.509-2.215"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.245 21.706c1.104.139 3.049-.425 3.622-3.377s-1.604-4.207-3.895-4.265s-5.7 2.302-3.105 6.535M8.88 32.586c2.34.086 4.96-2.363 6.507-1.804s3.995 3.071 3.995 3.071"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.366 32.972c.88.774 3.092 2.04 4.446.71c.494.944 1.095 3.908-2.986 3.414c-.107.86-1.332 1.096-2.62 1.117m6.866-6.021l-1.261 1.489M4.5 18.904a27.793 27.793 0 0 1 6.757.466m3.121-3.174a3.54 3.54 0 0 1 3.028 2.429"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.506 18.496a5.388 5.388 0 0 1 1.386-1.544"/>`),
		g.Group(children),
	)
}