package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseWithGarden(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#855c52" d="M44.46 97.08h23.16v30.89H44.46z"/><defs><path id="notoV1HouseWithGarden0" d="M41.99 48.17C38.86 40.8 30.95 28.86 24.2 28.86c-6.79 0-14.07 11.06-17.01 17.47C3.5 54.35.43 65.14.43 78.87c0 17.41 7.51 27.03 15.44 30.89h-.01l.09.04c.77.38 1.56.9 2.13 1.59c.66.82.73 2.54.73 3.12v7.63s-.09 2.91-.36 3.8c-.18.57-.78 1.27-.82 2.01h11.25v-38.2c0-3.03 1.62-6.7 3.86-8.73L47.2 67.86c-.98-7.72-2.93-14.29-5.21-19.69z"/></defs><use fill="#855c52" href="#notoV1HouseWithGarden0"/><clipPath id="notoV1HouseWithGarden1"><use href="#notoV1HouseWithGarden0"/></clipPath><path fill="#bdcf46" d="M15.95 109.81s5.1 1.93 10.1 1.93c0 0 5.44.52 9.94-2.36L50.13 67.6l-7.51-32.65l-16.88-9.88l-8.51.88L.98 47.84s-5.25 27.39-5.25 27.77c0 .37 4.38 26.14 4.38 26.14l7.88 5.38l7.96 2.68z" clip-path="url(#notoV1HouseWithGarden1)"/><path fill="#bdcf46" d="M123.78 113.84c-.49-.27-1.18-.46-1.52-.95c-.47-.69-.01-1.53-.01-2.28c0-3.37-2.29-6.61-5.67-7.3c-.63-.13-1.34.05-1.94-.22c-.87-.38-.75-1.25-.96-2.04c-.31-1.22-.9-2.35-1.78-3.26c-1.33-1.38-3.12-2.2-5.05-2.2c-1.15 0-2.23.31-3.2.83V128h16.8c3.54 0 6.59-2.85 7.03-6.3c.42-3.13-.88-6.31-3.7-7.86z"/><path fill="#fff" d="M71.5 91.41h22.79v24.69H71.5z"/><defs><path id="notoV1HouseWithGarden2" d="M68.75 58.47c-1.37-1.24-3.6-1.24-4.96 0L47.72 73.1L35.59 84.14c-1.36 1.24-2.48 3.76-2.48 5.61v34.87c0 1.84 1.51 3.35 3.35 3.35H49.4v-23c0-1.4 1.15-2.55 2.54-2.55h10.21c1.4 0 2.55 1.15 2.55 2.55v23h31.38c1.85 0 3.35-1.51 3.35-3.35V89.75c0-1.85-1.12-4.37-2.48-5.61l-28.2-25.67zm12.02 53.32h-3.64c-1.7 0-3.1-1.39-3.1-3.1v-3.12h6.74v6.22zm0-9.17h-6.74v-4.17c0-1.7 1.4-3.09 3.1-3.09h3.64v7.26zm9.7 6.07c0 1.7-1.39 3.1-3.1 3.1h-3.65v-6.21h6.74v3.11zm0-6.07h-6.74v-7.27h3.65c1.7 0 3.1 1.39 3.1 3.09v4.18z"/></defs><use fill="#fcc21b" href="#notoV1HouseWithGarden2"/><clipPath id="notoV1HouseWithGarden3"><use href="#notoV1HouseWithGarden2"/></clipPath><path fill="#ef6b00" d="M102.71 88.82H29.5l37.03-33.51z" clip-path="url(#notoV1HouseWithGarden3)"/>`),
		g.Group(children),
	)
}