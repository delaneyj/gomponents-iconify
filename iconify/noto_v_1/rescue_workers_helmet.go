package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RescueWorkersHelmet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><path id="notoV1RescueWorkersHelmet0" d="M0 0h128v128H0z"/></defs><clipPath id="notoV1RescueWorkersHelmet1"><use href="#notoV1RescueWorkersHelmet0"/></clipPath><path fill="#37b4e2" d="M64.71 11.75h.19h-.36h.17z" clip-path="url(#notoV1RescueWorkersHelmet1)"/><g clip-path="url(#notoV1RescueWorkersHelmet1)"><path fill="#875e54" d="M64.71 123.92c-31.32 0-56.79-25.3-56.79-56.4c0-31.1 25.48-56.4 56.79-56.4s56.79 25.3 56.79 56.4c0 31.1-25.47 56.4-56.79 56.4zm0-107.06c-28.42 0-51.55 22.73-51.55 50.67s23.12 50.66 51.55 50.66c28.42 0 51.55-22.73 51.55-50.66c0-27.94-23.12-50.67-51.55-50.67z"/><path fill="#875e54" d="M54.03 112.3h20.32v10.03H54.03z"/><path fill="#be1931" d="M125.41 60.07c-.03-.37-2.27.22-3.41-3.87c-3.98-14.23-15.58-48.67-57.99-48.67l-.01.95l-.01-.95c-42.53 0-52.97 33.14-56.94 47.64c-1.23 4.47-4.42 4.51-4.46 4.91c-.34 3.7 3.44 8.92 11.81 14.07c12.4 7.62 33.65 11.72 49.59 11.73h.02c15.94 0 37.19-4.11 49.59-11.73c8.37-5.15 12.15-10.37 11.81-14.08"/><path fill="#ea596e" d="m53.23 7.3l1.09-.32c6.3-1.84 13.01-1.73 19.25.32l13.16 56.29l-.3.06a111.764 111.764 0 0 1-44.81.01l-.35-.07L53.23 7.3z"/><path fill="#fff" d="M60.65 32.86h6.69v24.53h-6.69z"/><path fill="#fff" d="M50.62 41.79h26.76v7.81H50.62z"/><path fill="#be1931" d="M7.64 53.05s-4.58 2.74-5.05 7.02l3.87 1.59l1.18-8.61zm112.74-.72s4.25 2.42 5.05 7.78l-3.87 1.59l-1.18-9.37z"/><path fill="#ea596e" d="m125.43 60.4l-7.49 2.27C81.52 73.33 44.1 73.12 7.74 62.05l-5.17-1.56l.18-1.11a3.68 3.68 0 0 1 4.76-2.9l.23.07c36.36 11.07 73.79 11.28 110.2.62l2.69-.82c2.08-.63 4.24.71 4.6 2.85l.2 1.2z"/></g>`),
		g.Group(children),
	)
}