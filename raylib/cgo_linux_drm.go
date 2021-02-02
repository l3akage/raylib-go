// +build linux,arm,drm,!android

package rl

/*
#cgo linux,arm,drm LDFLAGS: -L/opt/vc/lib -L/opt/vc/lib64 -lGLESv2 -lEGL -lpthread -lrt -lm -lgbm -ldrm -ldl
#cgo linux,arm,drm CFLAGS: -D_DEFAULT_SOURCE -DPLATFORM_DRM -DEGL_NO_X11 -I/opt/vc/include -I/opt/vc/include/interface/vcos -I/opt/vc/include/interface/vmcs_host/linux -I/opt/vc/include/interface/vcos/pthreads -I/usr/include/libdrm

#cgo linux,arm,drm,!opengl11,!opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_ES2
*/
import "C"
