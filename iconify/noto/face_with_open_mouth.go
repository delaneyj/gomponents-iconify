package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithOpenMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<radialGradient id="notoFaceWithOpenMouth0" cx="63.6" cy="1696.9" r="56.96" gradientTransform="translate(0 -1634)" gradientUnits="userSpaceOnUse"><stop offset=".5" stop-color="#FDE030"/><stop offset=".92" stop-color="#F7C02B"/><stop offset="1" stop-color="#F4A223"/></radialGradient><path fill="url(#notoFaceWithOpenMouth0)" d="M63.6 118.8c-27.9 0-58-17.5-58-55.9S35.7 7 63.6 7c15.5 0 29.8 5.1 40.4 14.4c11.5 10.2 17.6 24.6 17.6 41.5s-6.1 31.2-17.6 41.4c-10.6 9.3-25 14.5-40.4 14.5z"/><path fill="#EB8F00" d="M111.49 29.67c5.33 8.6 8.11 18.84 8.11 30.23c0 16.9-6.1 31.2-17.6 41.4c-10.6 9.3-25 14.5-40.4 14.5c-18.06 0-37-7.35-48.18-22.94c10.76 17.66 31 25.94 50.18 25.94c15.4 0 29.8-5.2 40.4-14.5c11.5-10.2 17.6-24.5 17.6-41.4c0-12.74-3.47-24.06-10.11-33.23z"/><path fill="#422B0D" d="M84 87.92c0 8.16-9 13-20 13s-20-4.8-20-13s9-17 20-17s20 8.83 20 17zM44 40.94c-4.19 0-8 3.54-8 9.42s3.81 9.41 8 9.41c4.2 0 8-3.54 8-9.41s-3.76-9.42-8-9.42z"/><path fill="#896024" d="M43.65 44.87a2.874 2.874 0 0 0-3.82 1.34c-.53 1.11-.29 2.44.6 3.3c1.42.68 3.13.08 3.82-1.34c.53-1.11.29-2.44-.6-3.3z"/><path fill="#422B0D" d="M82.4 40.94c-4.19 0-8 3.54-8 9.42s3.81 9.41 8 9.41c4.19 0 8-3.54 8-9.41s-3.81-9.42-8-9.42z"/><path fill="#896024" d="M82 44.87a2.874 2.874 0 0 0-3.82 1.34c-.53 1.11-.29 2.44.6 3.3c1.42.68 3.13.08 3.82-1.34c.53-1.11.29-2.44-.6-3.3z"/>`),
		g.Group(children),
	)
}