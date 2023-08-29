package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PetFirstAid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.123 43.5c.564-15.254-5.578-20.398-10.426-23.59c1.182-5.795 5.427-8.926 5.427-8.926"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.098 10.331s-1.973 3.193-3.747 3.26c0 0-.577-5.832-.355-7.007C22.437 4.9 25.541 4.5 27.315 4.5s3.193 1.197 3.437 3.813c1.685.023 3.281.799 3.614 1.53c-.111 1.863-1.619 2.816-3.082 3.37s-2.66 1.464-2.66 2.35s.398 4.235 2.881 4.701s6.408 1.685 6.408 5.698s-1.353 3.348-1.353 7.28s.414 3.045.414 5.055s-.65 3.784-1.3 4.966"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.333 42.909s.74-3.43.74-5.055s-1.382-3.74-2.839-5.573c-1.256-1.581-4.533-.204-4.414 3.904m10.303-14.043c3.305.34 5.611-2.558 3.66-6.046M15.697 19.91c.207.916.377 2.749.377 2.749c-1.132 0-2.159.957-2.802 1.379c-1.312.86-2.686 1.148-2.638 2.7c.473 1.242 2.594 1.382 4.132 1.353l-.034.864c-1.531 1.464-4.423 4.746-4.423 8.78c0 4.7 3.437 5.528 5.388 5.528s11.68-4.918 8.489-13.51"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.708 30.138c.916.089 1.921.236 1.921 1.685S7.501 34.63 7.501 38.03s2.284 5.232 8.196 5.232m21.186-6.212c-.559-2.182-3.101-4.577-3.81-4.931s-.751 3.236-.236 4.448"/>`),
		g.Group(children),
	)
}