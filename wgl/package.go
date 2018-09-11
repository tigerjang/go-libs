// Copyright (c) 2010 Khronos Group.
// This material may be distributed subject to the terms and conditions
// set forth in the Open Publication License, v 1.0, 8 June 1999.
// http://opencontent.org/openpub/.
//
// Copyright (c) 1991-2006 Silicon Graphics, Inc.
// This document is licensed under the SGI Free Software B License.
// For details, see http://oss.sgi.com/projects/FreeB.

// Package wgl implements Go bindings to OpenGL.
//
// This package was automatically generated using Glow:
//  http://github.com/go-gl/glow
//
package wgl

// #cgo darwin        LDFLAGS: -framework OpenGL
// #cgo linux freebsd LDFLAGS: -lGL
// #cgo windows       LDFLAGS: -lopengl32
// #if defined(_WIN32) && !defined(APIENTRY) && !defined(__CYGWIN__) && !defined(__SCITECH_SNAP__)
// #ifndef WIN32_LEAN_AND_MEAN
// #define WIN32_LEAN_AND_MEAN 1
// #endif
// #include <windows.h>
// #endif
// #ifndef APIENTRY
// #define APIENTRY
// #endif
// #ifndef APIENTRYP
// #define APIENTRYP APIENTRY *
// #endif
// #ifndef GLAPI
// #define GLAPI extern
// #endif
// struct _GPU_DEVICE {
//     DWORD  cb;
//     CHAR   DeviceName[32];
//     CHAR   DeviceString[128];
//     DWORD  Flags;
//     RECT   rcVirtualScreen;
// };
// DECLARE_HANDLE(HPBUFFERARB);
// DECLARE_HANDLE(HPBUFFEREXT);
// DECLARE_HANDLE(HVIDEOOUTPUTDEVICENV);
// DECLARE_HANDLE(HPVIDEODEV);
// DECLARE_HANDLE(HGPUNV);
// DECLARE_HANDLE(HVIDEOINPUTDEVICENV);
// typedef struct _GPU_DEVICE *PGPU_DEVICE;
// typedef int  (APIENTRYP GPCHOOSEPIXELFORMAT)(HDC  hDc, const PIXELFORMATDESCRIPTOR * pPfd);
// typedef int  (APIENTRYP GPDESCRIBEPIXELFORMAT)(HDC  hdc, int  ipfd, UINT  cjpfd, const PIXELFORMATDESCRIPTOR * ppfd);
// typedef UINT  (APIENTRYP GPGETENHMETAFILEPIXELFORMAT)(HENHMETAFILE  hemf, const PIXELFORMATDESCRIPTOR * ppfd);
// typedef int  (APIENTRYP GPGETPIXELFORMAT)(HDC  hdc);
// typedef BOOL  (APIENTRYP GPSETPIXELFORMAT)(HDC  hdc, int  ipfd, const PIXELFORMATDESCRIPTOR * ppfd);
// typedef BOOL  (APIENTRYP GPSWAPBUFFERS)(HDC  hdc);
// typedef void * (APIENTRYP GPALLOCATEMEMORYNV)(GLsizei  size, GLfloat  readfreq, GLfloat  writefreq, GLfloat  priority);
// typedef BOOL  (APIENTRYP GPASSOCIATEIMAGEBUFFEREVENTSI3D)(HDC  hDC, const HANDLE * pEvent, const LPVOID * pAddress, const DWORD * pSize, UINT  count);
// typedef BOOL  (APIENTRYP GPBEGINFRAMETRACKINGI3D)();
// typedef GLboolean  (APIENTRYP GPBINDDISPLAYCOLORTABLEEXT)(GLushort  id);
// typedef BOOL  (APIENTRYP GPBINDSWAPBARRIERNV)(GLuint  group, GLuint  barrier);
// typedef BOOL  (APIENTRYP GPBINDTEXIMAGEARB)(HPBUFFERARB  hPbuffer, int  iBuffer);
// typedef BOOL  (APIENTRYP GPBINDVIDEOCAPTUREDEVICENV)(UINT  uVideoSlot, HVIDEOINPUTDEVICENV  hDevice);
// typedef BOOL  (APIENTRYP GPBINDVIDEODEVICENV)(HDC  hDC, unsigned int  uVideoSlot, HVIDEOOUTPUTDEVICENV  hVideoDevice, const int * piAttribList);
// typedef BOOL  (APIENTRYP GPBINDVIDEOIMAGENV)(HPVIDEODEV  hVideoDevice, HPBUFFERARB  hPbuffer, int  iVideoBuffer);
// typedef VOID  (APIENTRYP GPBLITCONTEXTFRAMEBUFFERAMD)(HGLRC  dstCtx, GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter);
// typedef BOOL  (APIENTRYP GPCHOOSEPIXELFORMATARB)(HDC  hdc, const int * piAttribIList, const FLOAT * pfAttribFList, UINT  nMaxFormats, int * piFormats, UINT * nNumFormats);
// typedef BOOL  (APIENTRYP GPCHOOSEPIXELFORMATEXT)(HDC  hdc, const int * piAttribIList, const FLOAT * pfAttribFList, UINT  nMaxFormats, int * piFormats, UINT * nNumFormats);
// typedef BOOL  (APIENTRYP GPCOPYCONTEXT)(HGLRC  hglrcSrc, HGLRC  hglrcDst, UINT  mask);
// typedef BOOL  (APIENTRYP GPCOPYIMAGESUBDATANV)(HGLRC  hSrcRC, GLuint  srcName, GLenum  srcTarget, GLint  srcLevel, GLint  srcX, GLint  srcY, GLint  srcZ, HGLRC  hDstRC, GLuint  dstName, GLenum  dstTarget, GLint  dstLevel, GLint  dstX, GLint  dstY, GLint  dstZ, GLsizei  width, GLsizei  height, GLsizei  depth);
// typedef HDC  (APIENTRYP GPCREATEAFFINITYDCNV)(const HGPUNV * phGpuList);
// typedef HGLRC  (APIENTRYP GPCREATEASSOCIATEDCONTEXTAMD)(UINT  id);
// typedef HGLRC  (APIENTRYP GPCREATEASSOCIATEDCONTEXTATTRIBSAMD)(UINT  id, HGLRC  hShareContext, const int * attribList);
// typedef HANDLE  (APIENTRYP GPCREATEBUFFERREGIONARB)(HDC  hDC, int  iLayerPlane, UINT  uType);
// typedef HGLRC  (APIENTRYP GPCREATECONTEXT)(HDC  hDc);
// typedef HGLRC  (APIENTRYP GPCREATECONTEXTATTRIBSARB)(HDC  hDC, HGLRC  hShareContext, const int * attribList);
// typedef GLboolean  (APIENTRYP GPCREATEDISPLAYCOLORTABLEEXT)(GLushort  id);
// typedef LPVOID  (APIENTRYP GPCREATEIMAGEBUFFERI3D)(HDC  hDC, DWORD  dwSize, UINT  uFlags);
// typedef HGLRC  (APIENTRYP GPCREATELAYERCONTEXT)(HDC  hDc, int  level);
// typedef HPBUFFERARB  (APIENTRYP GPCREATEPBUFFERARB)(HDC  hDC, int  iPixelFormat, int  iWidth, int  iHeight, const int * piAttribList);
// typedef HPBUFFEREXT  (APIENTRYP GPCREATEPBUFFEREXT)(HDC  hDC, int  iPixelFormat, int  iWidth, int  iHeight, const int * piAttribList);
// typedef BOOL  (APIENTRYP GPDXCLOSEDEVICENV)(HANDLE  hDevice);
// typedef BOOL  (APIENTRYP GPDXLOCKOBJECTSNV)(HANDLE  hDevice, GLint  count, HANDLE * hObjects);
// typedef BOOL  (APIENTRYP GPDXOBJECTACCESSNV)(HANDLE  hObject, GLenum  access);
// typedef HANDLE  (APIENTRYP GPDXOPENDEVICENV)(void * dxDevice);
// typedef HANDLE  (APIENTRYP GPDXREGISTEROBJECTNV)(HANDLE  hDevice, void * dxObject, GLuint  name, GLenum  type, GLenum  access);
// typedef BOOL  (APIENTRYP GPDXSETRESOURCESHAREHANDLENV)(void * dxObject, HANDLE  shareHandle);
// typedef BOOL  (APIENTRYP GPDXUNLOCKOBJECTSNV)(HANDLE  hDevice, GLint  count, HANDLE * hObjects);
// typedef BOOL  (APIENTRYP GPDXUNREGISTEROBJECTNV)(HANDLE  hDevice, HANDLE  hObject);
// typedef BOOL  (APIENTRYP GPDELAYBEFORESWAPNV)(HDC  hDC, GLfloat  seconds);
// typedef BOOL  (APIENTRYP GPDELETEASSOCIATEDCONTEXTAMD)(HGLRC  hglrc);
// typedef VOID  (APIENTRYP GPDELETEBUFFERREGIONARB)(HANDLE  hRegion);
// typedef BOOL  (APIENTRYP GPDELETECONTEXT)(HGLRC  oldContext);
// typedef BOOL  (APIENTRYP GPDELETEDCNV)(HDC  hdc);
// typedef BOOL  (APIENTRYP GPDESCRIBELAYERPLANE)(HDC  hDc, int  pixelFormat, int  layerPlane, UINT  nBytes, const LAYERPLANEDESCRIPTOR * plpd);
// typedef VOID  (APIENTRYP GPDESTROYDISPLAYCOLORTABLEEXT)(GLushort  id);
// typedef BOOL  (APIENTRYP GPDESTROYIMAGEBUFFERI3D)(HDC  hDC, LPVOID  pAddress);
// typedef BOOL  (APIENTRYP GPDESTROYPBUFFERARB)(HPBUFFERARB  hPbuffer);
// typedef BOOL  (APIENTRYP GPDESTROYPBUFFEREXT)(HPBUFFEREXT  hPbuffer);
// typedef BOOL  (APIENTRYP GPDISABLEFRAMELOCKI3D)();
// typedef BOOL  (APIENTRYP GPDISABLEGENLOCKI3D)(HDC  hDC);
// typedef BOOL  (APIENTRYP GPENABLEFRAMELOCKI3D)();
// typedef BOOL  (APIENTRYP GPENABLEGENLOCKI3D)(HDC  hDC);
// typedef BOOL  (APIENTRYP GPENDFRAMETRACKINGI3D)();
// typedef BOOL  (APIENTRYP GPENUMGPUDEVICESNV)(HGPUNV  hGpu, UINT  iDeviceIndex, PGPU_DEVICE  lpGpuDevice);
// typedef BOOL  (APIENTRYP GPENUMGPUSFROMAFFINITYDCNV)(HDC  hAffinityDC, UINT  iGpuIndex, HGPUNV * hGpu);
// typedef BOOL  (APIENTRYP GPENUMGPUSNV)(UINT  iGpuIndex, HGPUNV * phGpu);
// typedef UINT  (APIENTRYP GPENUMERATEVIDEOCAPTUREDEVICESNV)(HDC  hDc, HVIDEOINPUTDEVICENV * phDeviceList);
// typedef int  (APIENTRYP GPENUMERATEVIDEODEVICESNV)(HDC  hDC, HVIDEOOUTPUTDEVICENV * phDeviceList);
// typedef void  (APIENTRYP GPFREEMEMORYNV)(void * pointer);
// typedef BOOL  (APIENTRYP GPGENLOCKSAMPLERATEI3D)(HDC  hDC, UINT  uRate);
// typedef BOOL  (APIENTRYP GPGENLOCKSOURCEDELAYI3D)(HDC  hDC, UINT  uDelay);
// typedef BOOL  (APIENTRYP GPGENLOCKSOURCEEDGEI3D)(HDC  hDC, UINT  uEdge);
// typedef BOOL  (APIENTRYP GPGENLOCKSOURCEI3D)(HDC  hDC, UINT  uSource);
// typedef UINT  (APIENTRYP GPGETCONTEXTGPUIDAMD)(HGLRC  hglrc);
// typedef HGLRC  (APIENTRYP GPGETCURRENTASSOCIATEDCONTEXTAMD)();
// typedef HGLRC  (APIENTRYP GPGETCURRENTCONTEXT)();
// typedef HDC  (APIENTRYP GPGETCURRENTDC)();
// typedef HDC  (APIENTRYP GPGETCURRENTREADDCARB)();
// typedef HDC  (APIENTRYP GPGETCURRENTREADDCEXT)();
// typedef BOOL  (APIENTRYP GPGETDIGITALVIDEOPARAMETERSI3D)(HDC  hDC, int  iAttribute, int * piValue);
// typedef const char * (APIENTRYP GPGETEXTENSIONSSTRINGARB)(HDC  hdc);
// typedef const char * (APIENTRYP GPGETEXTENSIONSSTRINGEXT)();
// typedef BOOL  (APIENTRYP GPGETFRAMEUSAGEI3D)(float * pUsage);
// typedef UINT  (APIENTRYP GPGETGPUIDSAMD)(UINT  maxCount, UINT * ids);
// typedef INT  (APIENTRYP GPGETGPUINFOAMD)(UINT  id, int  property, GLenum  dataType, UINT  size, void * data);
// typedef BOOL  (APIENTRYP GPGETGAMMATABLEI3D)(HDC  hDC, int  iEntries, USHORT * puRed, USHORT * puGreen, USHORT * puBlue);
// typedef BOOL  (APIENTRYP GPGETGAMMATABLEPARAMETERSI3D)(HDC  hDC, int  iAttribute, int * piValue);
// typedef BOOL  (APIENTRYP GPGETGENLOCKSAMPLERATEI3D)(HDC  hDC, UINT * uRate);
// typedef BOOL  (APIENTRYP GPGETGENLOCKSOURCEDELAYI3D)(HDC  hDC, UINT * uDelay);
// typedef BOOL  (APIENTRYP GPGETGENLOCKSOURCEEDGEI3D)(HDC  hDC, UINT * uEdge);
// typedef BOOL  (APIENTRYP GPGETGENLOCKSOURCEI3D)(HDC  hDC, UINT * uSource);
// typedef int  (APIENTRYP GPGETLAYERPALETTEENTRIES)(HDC  hdc, int  iLayerPlane, int  iStart, int  cEntries, const COLORREF * pcr);
// typedef BOOL  (APIENTRYP GPGETMSCRATEOML)(HDC  hdc, INT32 * numerator, INT32 * denominator);
// typedef HDC  (APIENTRYP GPGETPBUFFERDCARB)(HPBUFFERARB  hPbuffer);
// typedef HDC  (APIENTRYP GPGETPBUFFERDCEXT)(HPBUFFEREXT  hPbuffer);
// typedef BOOL  (APIENTRYP GPGETPIXELFORMATATTRIBFVARB)(HDC  hdc, int  iPixelFormat, int  iLayerPlane, UINT  nAttributes, const int * piAttributes, FLOAT * pfValues);
// typedef BOOL  (APIENTRYP GPGETPIXELFORMATATTRIBFVEXT)(HDC  hdc, int  iPixelFormat, int  iLayerPlane, UINT  nAttributes, int * piAttributes, FLOAT * pfValues);
// typedef BOOL  (APIENTRYP GPGETPIXELFORMATATTRIBIVARB)(HDC  hdc, int  iPixelFormat, int  iLayerPlane, UINT  nAttributes, const int * piAttributes, int * piValues);
// typedef BOOL  (APIENTRYP GPGETPIXELFORMATATTRIBIVEXT)(HDC  hdc, int  iPixelFormat, int  iLayerPlane, UINT  nAttributes, int * piAttributes, int * piValues);
// typedef PROC  (APIENTRYP GPGETPROCADDRESS)(LPCSTR  lpszProc);
// typedef int  (APIENTRYP GPGETSWAPINTERVALEXT)();
// typedef BOOL  (APIENTRYP GPGETSYNCVALUESOML)(HDC  hdc, INT64 * ust, INT64 * msc, INT64 * sbc);
// typedef BOOL  (APIENTRYP GPGETVIDEODEVICENV)(HDC  hDC, int  numDevices, HPVIDEODEV * hVideoDevice);
// typedef BOOL  (APIENTRYP GPGETVIDEOINFONV)(HPVIDEODEV  hpVideoDevice, unsigned long * pulCounterOutputPbuffer, unsigned long * pulCounterOutputVideo);
// typedef BOOL  (APIENTRYP GPISENABLEDFRAMELOCKI3D)(BOOL * pFlag);
// typedef BOOL  (APIENTRYP GPISENABLEDGENLOCKI3D)(HDC  hDC, BOOL * pFlag);
// typedef BOOL  (APIENTRYP GPJOINSWAPGROUPNV)(HDC  hDC, GLuint  group);
// typedef GLboolean  (APIENTRYP GPLOADDISPLAYCOLORTABLEEXT)(const GLushort * table, GLuint  length);
// typedef BOOL  (APIENTRYP GPLOCKVIDEOCAPTUREDEVICENV)(HDC  hDc, HVIDEOINPUTDEVICENV  hDevice);
// typedef BOOL  (APIENTRYP GPMAKEASSOCIATEDCONTEXTCURRENTAMD)(HGLRC  hglrc);
// typedef BOOL  (APIENTRYP GPMAKECONTEXTCURRENTARB)(HDC  hDrawDC, HDC  hReadDC, HGLRC  hglrc);
// typedef BOOL  (APIENTRYP GPMAKECONTEXTCURRENTEXT)(HDC  hDrawDC, HDC  hReadDC, HGLRC  hglrc);
// typedef BOOL  (APIENTRYP GPMAKECURRENT)(HDC  hDc, HGLRC  newContext);
// typedef BOOL  (APIENTRYP GPQUERYCURRENTCONTEXTNV)(int  iAttribute, int * piValue);
// typedef BOOL  (APIENTRYP GPQUERYFRAMECOUNTNV)(HDC  hDC, GLuint * count);
// typedef BOOL  (APIENTRYP GPQUERYFRAMELOCKMASTERI3D)(BOOL * pFlag);
// typedef BOOL  (APIENTRYP GPQUERYFRAMETRACKINGI3D)(DWORD * pFrameCount, DWORD * pMissedFrames, float * pLastMissedUsage);
// typedef BOOL  (APIENTRYP GPQUERYGENLOCKMAXSOURCEDELAYI3D)(HDC  hDC, UINT * uMaxLineDelay, UINT * uMaxPixelDelay);
// typedef BOOL  (APIENTRYP GPQUERYMAXSWAPGROUPSNV)(HDC  hDC, GLuint * maxGroups, GLuint * maxBarriers);
// typedef BOOL  (APIENTRYP GPQUERYPBUFFERARB)(HPBUFFERARB  hPbuffer, int  iAttribute, int * piValue);
// typedef BOOL  (APIENTRYP GPQUERYPBUFFEREXT)(HPBUFFEREXT  hPbuffer, int  iAttribute, int * piValue);
// typedef BOOL  (APIENTRYP GPQUERYSWAPGROUPNV)(HDC  hDC, GLuint * group, GLuint * barrier);
// typedef BOOL  (APIENTRYP GPQUERYVIDEOCAPTUREDEVICENV)(HDC  hDc, HVIDEOINPUTDEVICENV  hDevice, int  iAttribute, int * piValue);
// typedef BOOL  (APIENTRYP GPREALIZELAYERPALETTE)(HDC  hdc, int  iLayerPlane, BOOL  bRealize);
// typedef BOOL  (APIENTRYP GPRELEASEIMAGEBUFFEREVENTSI3D)(HDC  hDC, const LPVOID * pAddress, UINT  count);
// typedef int  (APIENTRYP GPRELEASEPBUFFERDCARB)(HPBUFFERARB  hPbuffer, HDC  hDC);
// typedef int  (APIENTRYP GPRELEASEPBUFFERDCEXT)(HPBUFFEREXT  hPbuffer, HDC  hDC);
// typedef BOOL  (APIENTRYP GPRELEASETEXIMAGEARB)(HPBUFFERARB  hPbuffer, int  iBuffer);
// typedef BOOL  (APIENTRYP GPRELEASEVIDEOCAPTUREDEVICENV)(HDC  hDc, HVIDEOINPUTDEVICENV  hDevice);
// typedef BOOL  (APIENTRYP GPRELEASEVIDEODEVICENV)(HPVIDEODEV  hVideoDevice);
// typedef BOOL  (APIENTRYP GPRELEASEVIDEOIMAGENV)(HPBUFFERARB  hPbuffer, int  iVideoBuffer);
// typedef BOOL  (APIENTRYP GPRESETFRAMECOUNTNV)(HDC  hDC);
// typedef BOOL  (APIENTRYP GPRESTOREBUFFERREGIONARB)(HANDLE  hRegion, int  x, int  y, int  width, int  height, int  xSrc, int  ySrc);
// typedef BOOL  (APIENTRYP GPSAVEBUFFERREGIONARB)(HANDLE  hRegion, int  x, int  y, int  width, int  height);
// typedef BOOL  (APIENTRYP GPSENDPBUFFERTOVIDEONV)(HPBUFFERARB  hPbuffer, int  iBufferType, unsigned long * pulCounterPbuffer, BOOL  bBlock);
// typedef BOOL  (APIENTRYP GPSETDIGITALVIDEOPARAMETERSI3D)(HDC  hDC, int  iAttribute, const int * piValue);
// typedef BOOL  (APIENTRYP GPSETGAMMATABLEI3D)(HDC  hDC, int  iEntries, const USHORT * puRed, const USHORT * puGreen, const USHORT * puBlue);
// typedef BOOL  (APIENTRYP GPSETGAMMATABLEPARAMETERSI3D)(HDC  hDC, int  iAttribute, const int * piValue);
// typedef int  (APIENTRYP GPSETLAYERPALETTEENTRIES)(HDC  hdc, int  iLayerPlane, int  iStart, int  cEntries, const COLORREF * pcr);
// typedef BOOL  (APIENTRYP GPSETPBUFFERATTRIBARB)(HPBUFFERARB  hPbuffer, const int * piAttribList);
// typedef BOOL  (APIENTRYP GPSETSTEREOEMITTERSTATE3DL)(HDC  hDC, UINT  uState);
// typedef BOOL  (APIENTRYP GPSHARELISTS)(HGLRC  hrcSrvShare, HGLRC  hrcSrvSource);
// typedef INT64  (APIENTRYP GPSWAPBUFFERSMSCOML)(HDC  hdc, INT64  target_msc, INT64  divisor, INT64  remainder);
// typedef BOOL  (APIENTRYP GPSWAPINTERVALEXT)(int  interval);
// typedef BOOL  (APIENTRYP GPSWAPLAYERBUFFERS)(HDC  hdc, UINT  fuFlags);
// typedef INT64  (APIENTRYP GPSWAPLAYERBUFFERSMSCOML)(HDC  hdc, int  fuPlanes, INT64  target_msc, INT64  divisor, INT64  remainder);
// typedef BOOL  (APIENTRYP GPUSEFONTBITMAPS)(HDC  hDC, DWORD  first, DWORD  count, DWORD  listBase);
// typedef BOOL  (APIENTRYP GPUSEFONTBITMAPSA)(HDC  hDC, DWORD  first, DWORD  count, DWORD  listBase);
// typedef BOOL  (APIENTRYP GPUSEFONTBITMAPSW)(HDC  hDC, DWORD  first, DWORD  count, DWORD  listBase);
// typedef BOOL  (APIENTRYP GPUSEFONTOUTLINES)(HDC  hDC, DWORD  first, DWORD  count, DWORD  listBase, FLOAT  deviation, FLOAT  extrusion, int  format, LPGLYPHMETRICSFLOAT  lpgmf);
// typedef BOOL  (APIENTRYP GPUSEFONTOUTLINESA)(HDC  hDC, DWORD  first, DWORD  count, DWORD  listBase, FLOAT  deviation, FLOAT  extrusion, int  format, LPGLYPHMETRICSFLOAT  lpgmf);
// typedef BOOL  (APIENTRYP GPUSEFONTOUTLINESW)(HDC  hDC, DWORD  first, DWORD  count, DWORD  listBase, FLOAT  deviation, FLOAT  extrusion, int  format, LPGLYPHMETRICSFLOAT  lpgmf);
// typedef BOOL  (APIENTRYP GPWAITFORMSCOML)(HDC  hdc, INT64  target_msc, INT64  divisor, INT64  remainder, INT64 * ust, INT64 * msc, INT64 * sbc);
// typedef BOOL  (APIENTRYP GPWAITFORSBCOML)(HDC  hdc, INT64  target_sbc, INT64 * ust, INT64 * msc, INT64 * sbc);
// static int  glowChoosePixelFormat(GPCHOOSEPIXELFORMAT fnptr, HDC  hDc, const PIXELFORMATDESCRIPTOR * pPfd) {
//   return (*fnptr)(hDc, pPfd);
// }
// static int  glowDescribePixelFormat(GPDESCRIBEPIXELFORMAT fnptr, HDC  hdc, int  ipfd, UINT  cjpfd, const PIXELFORMATDESCRIPTOR * ppfd) {
//   return (*fnptr)(hdc, ipfd, cjpfd, ppfd);
// }
// static UINT  glowGetEnhMetaFilePixelFormat(GPGETENHMETAFILEPIXELFORMAT fnptr, HENHMETAFILE  hemf, const PIXELFORMATDESCRIPTOR * ppfd) {
//   return (*fnptr)(hemf, ppfd);
// }
// static int  glowGetPixelFormat(GPGETPIXELFORMAT fnptr, HDC  hdc) {
//   return (*fnptr)(hdc);
// }
// static BOOL  glowSetPixelFormat(GPSETPIXELFORMAT fnptr, HDC  hdc, int  ipfd, const PIXELFORMATDESCRIPTOR * ppfd) {
//   return (*fnptr)(hdc, ipfd, ppfd);
// }
// static BOOL  glowSwapBuffers(GPSWAPBUFFERS fnptr, HDC  hdc) {
//   return (*fnptr)(hdc);
// }
// static void * glowAllocateMemoryNV(GPALLOCATEMEMORYNV fnptr, GLsizei  size, GLfloat  readfreq, GLfloat  writefreq, GLfloat  priority) {
//   return (*fnptr)(size, readfreq, writefreq, priority);
// }
// static BOOL  glowAssociateImageBufferEventsI3D(GPASSOCIATEIMAGEBUFFEREVENTSI3D fnptr, HDC  hDC, const HANDLE * pEvent, const LPVOID * pAddress, const DWORD * pSize, UINT  count) {
//   return (*fnptr)(hDC, pEvent, pAddress, pSize, count);
// }
// static BOOL  glowBeginFrameTrackingI3D(GPBEGINFRAMETRACKINGI3D fnptr) {
//   return (*fnptr)();
// }
// static GLboolean  glowBindDisplayColorTableEXT(GPBINDDISPLAYCOLORTABLEEXT fnptr, GLushort  id) {
//   return (*fnptr)(id);
// }
// static BOOL  glowBindSwapBarrierNV(GPBINDSWAPBARRIERNV fnptr, GLuint  group, GLuint  barrier) {
//   return (*fnptr)(group, barrier);
// }
// static BOOL  glowBindTexImageARB(GPBINDTEXIMAGEARB fnptr, HPBUFFERARB  hPbuffer, int  iBuffer) {
//   return (*fnptr)(hPbuffer, iBuffer);
// }
// static BOOL  glowBindVideoCaptureDeviceNV(GPBINDVIDEOCAPTUREDEVICENV fnptr, UINT  uVideoSlot, HVIDEOINPUTDEVICENV  hDevice) {
//   return (*fnptr)(uVideoSlot, hDevice);
// }
// static BOOL  glowBindVideoDeviceNV(GPBINDVIDEODEVICENV fnptr, HDC  hDC, unsigned int  uVideoSlot, HVIDEOOUTPUTDEVICENV  hVideoDevice, const int * piAttribList) {
//   return (*fnptr)(hDC, uVideoSlot, hVideoDevice, piAttribList);
// }
// static BOOL  glowBindVideoImageNV(GPBINDVIDEOIMAGENV fnptr, HPVIDEODEV  hVideoDevice, HPBUFFERARB  hPbuffer, int  iVideoBuffer) {
//   return (*fnptr)(hVideoDevice, hPbuffer, iVideoBuffer);
// }
// static VOID  glowBlitContextFramebufferAMD(GPBLITCONTEXTFRAMEBUFFERAMD fnptr, HGLRC  dstCtx, GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter) {
//   return (*fnptr)(dstCtx, srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
// }
// static BOOL  glowChoosePixelFormatARB(GPCHOOSEPIXELFORMATARB fnptr, HDC  hdc, const int * piAttribIList, const FLOAT * pfAttribFList, UINT  nMaxFormats, int * piFormats, UINT * nNumFormats) {
//   return (*fnptr)(hdc, piAttribIList, pfAttribFList, nMaxFormats, piFormats, nNumFormats);
// }
// static BOOL  glowChoosePixelFormatEXT(GPCHOOSEPIXELFORMATEXT fnptr, HDC  hdc, const int * piAttribIList, const FLOAT * pfAttribFList, UINT  nMaxFormats, int * piFormats, UINT * nNumFormats) {
//   return (*fnptr)(hdc, piAttribIList, pfAttribFList, nMaxFormats, piFormats, nNumFormats);
// }
// static BOOL  glowCopyContext(GPCOPYCONTEXT fnptr, HGLRC  hglrcSrc, HGLRC  hglrcDst, UINT  mask) {
//   return (*fnptr)(hglrcSrc, hglrcDst, mask);
// }
// static BOOL  glowCopyImageSubDataNV(GPCOPYIMAGESUBDATANV fnptr, HGLRC  hSrcRC, GLuint  srcName, GLenum  srcTarget, GLint  srcLevel, GLint  srcX, GLint  srcY, GLint  srcZ, HGLRC  hDstRC, GLuint  dstName, GLenum  dstTarget, GLint  dstLevel, GLint  dstX, GLint  dstY, GLint  dstZ, GLsizei  width, GLsizei  height, GLsizei  depth) {
//   return (*fnptr)(hSrcRC, srcName, srcTarget, srcLevel, srcX, srcY, srcZ, hDstRC, dstName, dstTarget, dstLevel, dstX, dstY, dstZ, width, height, depth);
// }
// static HDC  glowCreateAffinityDCNV(GPCREATEAFFINITYDCNV fnptr, const HGPUNV * phGpuList) {
//   return (*fnptr)(phGpuList);
// }
// static HGLRC  glowCreateAssociatedContextAMD(GPCREATEASSOCIATEDCONTEXTAMD fnptr, UINT  id) {
//   return (*fnptr)(id);
// }
// static HGLRC  glowCreateAssociatedContextAttribsAMD(GPCREATEASSOCIATEDCONTEXTATTRIBSAMD fnptr, UINT  id, HGLRC  hShareContext, const int * attribList) {
//   return (*fnptr)(id, hShareContext, attribList);
// }
// static HANDLE  glowCreateBufferRegionARB(GPCREATEBUFFERREGIONARB fnptr, HDC  hDC, int  iLayerPlane, UINT  uType) {
//   return (*fnptr)(hDC, iLayerPlane, uType);
// }
// static HGLRC  glowCreateContext(GPCREATECONTEXT fnptr, HDC  hDc) {
//   return (*fnptr)(hDc);
// }
// static HGLRC  glowCreateContextAttribsARB(GPCREATECONTEXTATTRIBSARB fnptr, HDC  hDC, HGLRC  hShareContext, const int * attribList) {
//   return (*fnptr)(hDC, hShareContext, attribList);
// }
// static GLboolean  glowCreateDisplayColorTableEXT(GPCREATEDISPLAYCOLORTABLEEXT fnptr, GLushort  id) {
//   return (*fnptr)(id);
// }
// static LPVOID  glowCreateImageBufferI3D(GPCREATEIMAGEBUFFERI3D fnptr, HDC  hDC, DWORD  dwSize, UINT  uFlags) {
//   return (*fnptr)(hDC, dwSize, uFlags);
// }
// static HGLRC  glowCreateLayerContext(GPCREATELAYERCONTEXT fnptr, HDC  hDc, int  level) {
//   return (*fnptr)(hDc, level);
// }
// static HPBUFFERARB  glowCreatePbufferARB(GPCREATEPBUFFERARB fnptr, HDC  hDC, int  iPixelFormat, int  iWidth, int  iHeight, const int * piAttribList) {
//   return (*fnptr)(hDC, iPixelFormat, iWidth, iHeight, piAttribList);
// }
// static HPBUFFEREXT  glowCreatePbufferEXT(GPCREATEPBUFFEREXT fnptr, HDC  hDC, int  iPixelFormat, int  iWidth, int  iHeight, const int * piAttribList) {
//   return (*fnptr)(hDC, iPixelFormat, iWidth, iHeight, piAttribList);
// }
// static BOOL  glowDXCloseDeviceNV(GPDXCLOSEDEVICENV fnptr, HANDLE  hDevice) {
//   return (*fnptr)(hDevice);
// }
// static BOOL  glowDXLockObjectsNV(GPDXLOCKOBJECTSNV fnptr, HANDLE  hDevice, GLint  count, HANDLE * hObjects) {
//   return (*fnptr)(hDevice, count, hObjects);
// }
// static BOOL  glowDXObjectAccessNV(GPDXOBJECTACCESSNV fnptr, HANDLE  hObject, GLenum  access) {
//   return (*fnptr)(hObject, access);
// }
// static HANDLE  glowDXOpenDeviceNV(GPDXOPENDEVICENV fnptr, void * dxDevice) {
//   return (*fnptr)(dxDevice);
// }
// static HANDLE  glowDXRegisterObjectNV(GPDXREGISTEROBJECTNV fnptr, HANDLE  hDevice, void * dxObject, GLuint  name, GLenum  type, GLenum  access) {
//   return (*fnptr)(hDevice, dxObject, name, type, access);
// }
// static BOOL  glowDXSetResourceShareHandleNV(GPDXSETRESOURCESHAREHANDLENV fnptr, void * dxObject, HANDLE  shareHandle) {
//   return (*fnptr)(dxObject, shareHandle);
// }
// static BOOL  glowDXUnlockObjectsNV(GPDXUNLOCKOBJECTSNV fnptr, HANDLE  hDevice, GLint  count, HANDLE * hObjects) {
//   return (*fnptr)(hDevice, count, hObjects);
// }
// static BOOL  glowDXUnregisterObjectNV(GPDXUNREGISTEROBJECTNV fnptr, HANDLE  hDevice, HANDLE  hObject) {
//   return (*fnptr)(hDevice, hObject);
// }
// static BOOL  glowDelayBeforeSwapNV(GPDELAYBEFORESWAPNV fnptr, HDC  hDC, GLfloat  seconds) {
//   return (*fnptr)(hDC, seconds);
// }
// static BOOL  glowDeleteAssociatedContextAMD(GPDELETEASSOCIATEDCONTEXTAMD fnptr, HGLRC  hglrc) {
//   return (*fnptr)(hglrc);
// }
// static VOID  glowDeleteBufferRegionARB(GPDELETEBUFFERREGIONARB fnptr, HANDLE  hRegion) {
//   return (*fnptr)(hRegion);
// }
// static BOOL  glowDeleteContext(GPDELETECONTEXT fnptr, HGLRC  oldContext) {
//   return (*fnptr)(oldContext);
// }
// static BOOL  glowDeleteDCNV(GPDELETEDCNV fnptr, HDC  hdc) {
//   return (*fnptr)(hdc);
// }
// static BOOL  glowDescribeLayerPlane(GPDESCRIBELAYERPLANE fnptr, HDC  hDc, int  pixelFormat, int  layerPlane, UINT  nBytes, const LAYERPLANEDESCRIPTOR * plpd) {
//   return (*fnptr)(hDc, pixelFormat, layerPlane, nBytes, plpd);
// }
// static VOID  glowDestroyDisplayColorTableEXT(GPDESTROYDISPLAYCOLORTABLEEXT fnptr, GLushort  id) {
//   return (*fnptr)(id);
// }
// static BOOL  glowDestroyImageBufferI3D(GPDESTROYIMAGEBUFFERI3D fnptr, HDC  hDC, LPVOID  pAddress) {
//   return (*fnptr)(hDC, pAddress);
// }
// static BOOL  glowDestroyPbufferARB(GPDESTROYPBUFFERARB fnptr, HPBUFFERARB  hPbuffer) {
//   return (*fnptr)(hPbuffer);
// }
// static BOOL  glowDestroyPbufferEXT(GPDESTROYPBUFFEREXT fnptr, HPBUFFEREXT  hPbuffer) {
//   return (*fnptr)(hPbuffer);
// }
// static BOOL  glowDisableFrameLockI3D(GPDISABLEFRAMELOCKI3D fnptr) {
//   return (*fnptr)();
// }
// static BOOL  glowDisableGenlockI3D(GPDISABLEGENLOCKI3D fnptr, HDC  hDC) {
//   return (*fnptr)(hDC);
// }
// static BOOL  glowEnableFrameLockI3D(GPENABLEFRAMELOCKI3D fnptr) {
//   return (*fnptr)();
// }
// static BOOL  glowEnableGenlockI3D(GPENABLEGENLOCKI3D fnptr, HDC  hDC) {
//   return (*fnptr)(hDC);
// }
// static BOOL  glowEndFrameTrackingI3D(GPENDFRAMETRACKINGI3D fnptr) {
//   return (*fnptr)();
// }
// static BOOL  glowEnumGpuDevicesNV(GPENUMGPUDEVICESNV fnptr, HGPUNV  hGpu, UINT  iDeviceIndex, PGPU_DEVICE  lpGpuDevice) {
//   return (*fnptr)(hGpu, iDeviceIndex, lpGpuDevice);
// }
// static BOOL  glowEnumGpusFromAffinityDCNV(GPENUMGPUSFROMAFFINITYDCNV fnptr, HDC  hAffinityDC, UINT  iGpuIndex, HGPUNV * hGpu) {
//   return (*fnptr)(hAffinityDC, iGpuIndex, hGpu);
// }
// static BOOL  glowEnumGpusNV(GPENUMGPUSNV fnptr, UINT  iGpuIndex, HGPUNV * phGpu) {
//   return (*fnptr)(iGpuIndex, phGpu);
// }
// static UINT  glowEnumerateVideoCaptureDevicesNV(GPENUMERATEVIDEOCAPTUREDEVICESNV fnptr, HDC  hDc, HVIDEOINPUTDEVICENV * phDeviceList) {
//   return (*fnptr)(hDc, phDeviceList);
// }
// static int  glowEnumerateVideoDevicesNV(GPENUMERATEVIDEODEVICESNV fnptr, HDC  hDC, HVIDEOOUTPUTDEVICENV * phDeviceList) {
//   return (*fnptr)(hDC, phDeviceList);
// }
// static void  glowFreeMemoryNV(GPFREEMEMORYNV fnptr, void * pointer) {
//   (*fnptr)(pointer);
// }
// static BOOL  glowGenlockSampleRateI3D(GPGENLOCKSAMPLERATEI3D fnptr, HDC  hDC, UINT  uRate) {
//   return (*fnptr)(hDC, uRate);
// }
// static BOOL  glowGenlockSourceDelayI3D(GPGENLOCKSOURCEDELAYI3D fnptr, HDC  hDC, UINT  uDelay) {
//   return (*fnptr)(hDC, uDelay);
// }
// static BOOL  glowGenlockSourceEdgeI3D(GPGENLOCKSOURCEEDGEI3D fnptr, HDC  hDC, UINT  uEdge) {
//   return (*fnptr)(hDC, uEdge);
// }
// static BOOL  glowGenlockSourceI3D(GPGENLOCKSOURCEI3D fnptr, HDC  hDC, UINT  uSource) {
//   return (*fnptr)(hDC, uSource);
// }
// static UINT  glowGetContextGPUIDAMD(GPGETCONTEXTGPUIDAMD fnptr, HGLRC  hglrc) {
//   return (*fnptr)(hglrc);
// }
// static HGLRC  glowGetCurrentAssociatedContextAMD(GPGETCURRENTASSOCIATEDCONTEXTAMD fnptr) {
//   return (*fnptr)();
// }
// static HGLRC  glowGetCurrentContext(GPGETCURRENTCONTEXT fnptr) {
//   return (*fnptr)();
// }
// static HDC  glowGetCurrentDC(GPGETCURRENTDC fnptr) {
//   return (*fnptr)();
// }
// static HDC  glowGetCurrentReadDCARB(GPGETCURRENTREADDCARB fnptr) {
//   return (*fnptr)();
// }
// static HDC  glowGetCurrentReadDCEXT(GPGETCURRENTREADDCEXT fnptr) {
//   return (*fnptr)();
// }
// static BOOL  glowGetDigitalVideoParametersI3D(GPGETDIGITALVIDEOPARAMETERSI3D fnptr, HDC  hDC, int  iAttribute, int * piValue) {
//   return (*fnptr)(hDC, iAttribute, piValue);
// }
// static const char * glowGetExtensionsStringARB(GPGETEXTENSIONSSTRINGARB fnptr, HDC  hdc) {
//   return (*fnptr)(hdc);
// }
// static const char * glowGetExtensionsStringEXT(GPGETEXTENSIONSSTRINGEXT fnptr) {
//   return (*fnptr)();
// }
// static BOOL  glowGetFrameUsageI3D(GPGETFRAMEUSAGEI3D fnptr, float * pUsage) {
//   return (*fnptr)(pUsage);
// }
// static UINT  glowGetGPUIDsAMD(GPGETGPUIDSAMD fnptr, UINT  maxCount, UINT * ids) {
//   return (*fnptr)(maxCount, ids);
// }
// static INT  glowGetGPUInfoAMD(GPGETGPUINFOAMD fnptr, UINT  id, int  property, GLenum  dataType, UINT  size, void * data) {
//   return (*fnptr)(id, property, dataType, size, data);
// }
// static BOOL  glowGetGammaTableI3D(GPGETGAMMATABLEI3D fnptr, HDC  hDC, int  iEntries, USHORT * puRed, USHORT * puGreen, USHORT * puBlue) {
//   return (*fnptr)(hDC, iEntries, puRed, puGreen, puBlue);
// }
// static BOOL  glowGetGammaTableParametersI3D(GPGETGAMMATABLEPARAMETERSI3D fnptr, HDC  hDC, int  iAttribute, int * piValue) {
//   return (*fnptr)(hDC, iAttribute, piValue);
// }
// static BOOL  glowGetGenlockSampleRateI3D(GPGETGENLOCKSAMPLERATEI3D fnptr, HDC  hDC, UINT * uRate) {
//   return (*fnptr)(hDC, uRate);
// }
// static BOOL  glowGetGenlockSourceDelayI3D(GPGETGENLOCKSOURCEDELAYI3D fnptr, HDC  hDC, UINT * uDelay) {
//   return (*fnptr)(hDC, uDelay);
// }
// static BOOL  glowGetGenlockSourceEdgeI3D(GPGETGENLOCKSOURCEEDGEI3D fnptr, HDC  hDC, UINT * uEdge) {
//   return (*fnptr)(hDC, uEdge);
// }
// static BOOL  glowGetGenlockSourceI3D(GPGETGENLOCKSOURCEI3D fnptr, HDC  hDC, UINT * uSource) {
//   return (*fnptr)(hDC, uSource);
// }
// static int  glowGetLayerPaletteEntries(GPGETLAYERPALETTEENTRIES fnptr, HDC  hdc, int  iLayerPlane, int  iStart, int  cEntries, const COLORREF * pcr) {
//   return (*fnptr)(hdc, iLayerPlane, iStart, cEntries, pcr);
// }
// static BOOL  glowGetMscRateOML(GPGETMSCRATEOML fnptr, HDC  hdc, INT32 * numerator, INT32 * denominator) {
//   return (*fnptr)(hdc, numerator, denominator);
// }
// static HDC  glowGetPbufferDCARB(GPGETPBUFFERDCARB fnptr, HPBUFFERARB  hPbuffer) {
//   return (*fnptr)(hPbuffer);
// }
// static HDC  glowGetPbufferDCEXT(GPGETPBUFFERDCEXT fnptr, HPBUFFEREXT  hPbuffer) {
//   return (*fnptr)(hPbuffer);
// }
// static BOOL  glowGetPixelFormatAttribfvARB(GPGETPIXELFORMATATTRIBFVARB fnptr, HDC  hdc, int  iPixelFormat, int  iLayerPlane, UINT  nAttributes, const int * piAttributes, FLOAT * pfValues) {
//   return (*fnptr)(hdc, iPixelFormat, iLayerPlane, nAttributes, piAttributes, pfValues);
// }
// static BOOL  glowGetPixelFormatAttribfvEXT(GPGETPIXELFORMATATTRIBFVEXT fnptr, HDC  hdc, int  iPixelFormat, int  iLayerPlane, UINT  nAttributes, int * piAttributes, FLOAT * pfValues) {
//   return (*fnptr)(hdc, iPixelFormat, iLayerPlane, nAttributes, piAttributes, pfValues);
// }
// static BOOL  glowGetPixelFormatAttribivARB(GPGETPIXELFORMATATTRIBIVARB fnptr, HDC  hdc, int  iPixelFormat, int  iLayerPlane, UINT  nAttributes, const int * piAttributes, int * piValues) {
//   return (*fnptr)(hdc, iPixelFormat, iLayerPlane, nAttributes, piAttributes, piValues);
// }
// static BOOL  glowGetPixelFormatAttribivEXT(GPGETPIXELFORMATATTRIBIVEXT fnptr, HDC  hdc, int  iPixelFormat, int  iLayerPlane, UINT  nAttributes, int * piAttributes, int * piValues) {
//   return (*fnptr)(hdc, iPixelFormat, iLayerPlane, nAttributes, piAttributes, piValues);
// }
// static PROC  glowGetProcAddress(GPGETPROCADDRESS fnptr, LPCSTR  lpszProc) {
//   return (*fnptr)(lpszProc);
// }
// static int  glowGetSwapIntervalEXT(GPGETSWAPINTERVALEXT fnptr) {
//   return (*fnptr)();
// }
// static BOOL  glowGetSyncValuesOML(GPGETSYNCVALUESOML fnptr, HDC  hdc, INT64 * ust, INT64 * msc, INT64 * sbc) {
//   return (*fnptr)(hdc, ust, msc, sbc);
// }
// static BOOL  glowGetVideoDeviceNV(GPGETVIDEODEVICENV fnptr, HDC  hDC, int  numDevices, HPVIDEODEV * hVideoDevice) {
//   return (*fnptr)(hDC, numDevices, hVideoDevice);
// }
// static BOOL  glowGetVideoInfoNV(GPGETVIDEOINFONV fnptr, HPVIDEODEV  hpVideoDevice, unsigned long * pulCounterOutputPbuffer, unsigned long * pulCounterOutputVideo) {
//   return (*fnptr)(hpVideoDevice, pulCounterOutputPbuffer, pulCounterOutputVideo);
// }
// static BOOL  glowIsEnabledFrameLockI3D(GPISENABLEDFRAMELOCKI3D fnptr, BOOL * pFlag) {
//   return (*fnptr)(pFlag);
// }
// static BOOL  glowIsEnabledGenlockI3D(GPISENABLEDGENLOCKI3D fnptr, HDC  hDC, BOOL * pFlag) {
//   return (*fnptr)(hDC, pFlag);
// }
// static BOOL  glowJoinSwapGroupNV(GPJOINSWAPGROUPNV fnptr, HDC  hDC, GLuint  group) {
//   return (*fnptr)(hDC, group);
// }
// static GLboolean  glowLoadDisplayColorTableEXT(GPLOADDISPLAYCOLORTABLEEXT fnptr, const GLushort * table, GLuint  length) {
//   return (*fnptr)(table, length);
// }
// static BOOL  glowLockVideoCaptureDeviceNV(GPLOCKVIDEOCAPTUREDEVICENV fnptr, HDC  hDc, HVIDEOINPUTDEVICENV  hDevice) {
//   return (*fnptr)(hDc, hDevice);
// }
// static BOOL  glowMakeAssociatedContextCurrentAMD(GPMAKEASSOCIATEDCONTEXTCURRENTAMD fnptr, HGLRC  hglrc) {
//   return (*fnptr)(hglrc);
// }
// static BOOL  glowMakeContextCurrentARB(GPMAKECONTEXTCURRENTARB fnptr, HDC  hDrawDC, HDC  hReadDC, HGLRC  hglrc) {
//   return (*fnptr)(hDrawDC, hReadDC, hglrc);
// }
// static BOOL  glowMakeContextCurrentEXT(GPMAKECONTEXTCURRENTEXT fnptr, HDC  hDrawDC, HDC  hReadDC, HGLRC  hglrc) {
//   return (*fnptr)(hDrawDC, hReadDC, hglrc);
// }
// static BOOL  glowMakeCurrent(GPMAKECURRENT fnptr, HDC  hDc, HGLRC  newContext) {
//   return (*fnptr)(hDc, newContext);
// }
// static BOOL  glowQueryCurrentContextNV(GPQUERYCURRENTCONTEXTNV fnptr, int  iAttribute, int * piValue) {
//   return (*fnptr)(iAttribute, piValue);
// }
// static BOOL  glowQueryFrameCountNV(GPQUERYFRAMECOUNTNV fnptr, HDC  hDC, GLuint * count) {
//   return (*fnptr)(hDC, count);
// }
// static BOOL  glowQueryFrameLockMasterI3D(GPQUERYFRAMELOCKMASTERI3D fnptr, BOOL * pFlag) {
//   return (*fnptr)(pFlag);
// }
// static BOOL  glowQueryFrameTrackingI3D(GPQUERYFRAMETRACKINGI3D fnptr, DWORD * pFrameCount, DWORD * pMissedFrames, float * pLastMissedUsage) {
//   return (*fnptr)(pFrameCount, pMissedFrames, pLastMissedUsage);
// }
// static BOOL  glowQueryGenlockMaxSourceDelayI3D(GPQUERYGENLOCKMAXSOURCEDELAYI3D fnptr, HDC  hDC, UINT * uMaxLineDelay, UINT * uMaxPixelDelay) {
//   return (*fnptr)(hDC, uMaxLineDelay, uMaxPixelDelay);
// }
// static BOOL  glowQueryMaxSwapGroupsNV(GPQUERYMAXSWAPGROUPSNV fnptr, HDC  hDC, GLuint * maxGroups, GLuint * maxBarriers) {
//   return (*fnptr)(hDC, maxGroups, maxBarriers);
// }
// static BOOL  glowQueryPbufferARB(GPQUERYPBUFFERARB fnptr, HPBUFFERARB  hPbuffer, int  iAttribute, int * piValue) {
//   return (*fnptr)(hPbuffer, iAttribute, piValue);
// }
// static BOOL  glowQueryPbufferEXT(GPQUERYPBUFFEREXT fnptr, HPBUFFEREXT  hPbuffer, int  iAttribute, int * piValue) {
//   return (*fnptr)(hPbuffer, iAttribute, piValue);
// }
// static BOOL  glowQuerySwapGroupNV(GPQUERYSWAPGROUPNV fnptr, HDC  hDC, GLuint * group, GLuint * barrier) {
//   return (*fnptr)(hDC, group, barrier);
// }
// static BOOL  glowQueryVideoCaptureDeviceNV(GPQUERYVIDEOCAPTUREDEVICENV fnptr, HDC  hDc, HVIDEOINPUTDEVICENV  hDevice, int  iAttribute, int * piValue) {
//   return (*fnptr)(hDc, hDevice, iAttribute, piValue);
// }
// static BOOL  glowRealizeLayerPalette(GPREALIZELAYERPALETTE fnptr, HDC  hdc, int  iLayerPlane, BOOL  bRealize) {
//   return (*fnptr)(hdc, iLayerPlane, bRealize);
// }
// static BOOL  glowReleaseImageBufferEventsI3D(GPRELEASEIMAGEBUFFEREVENTSI3D fnptr, HDC  hDC, const LPVOID * pAddress, UINT  count) {
//   return (*fnptr)(hDC, pAddress, count);
// }
// static int  glowReleasePbufferDCARB(GPRELEASEPBUFFERDCARB fnptr, HPBUFFERARB  hPbuffer, HDC  hDC) {
//   return (*fnptr)(hPbuffer, hDC);
// }
// static int  glowReleasePbufferDCEXT(GPRELEASEPBUFFERDCEXT fnptr, HPBUFFEREXT  hPbuffer, HDC  hDC) {
//   return (*fnptr)(hPbuffer, hDC);
// }
// static BOOL  glowReleaseTexImageARB(GPRELEASETEXIMAGEARB fnptr, HPBUFFERARB  hPbuffer, int  iBuffer) {
//   return (*fnptr)(hPbuffer, iBuffer);
// }
// static BOOL  glowReleaseVideoCaptureDeviceNV(GPRELEASEVIDEOCAPTUREDEVICENV fnptr, HDC  hDc, HVIDEOINPUTDEVICENV  hDevice) {
//   return (*fnptr)(hDc, hDevice);
// }
// static BOOL  glowReleaseVideoDeviceNV(GPRELEASEVIDEODEVICENV fnptr, HPVIDEODEV  hVideoDevice) {
//   return (*fnptr)(hVideoDevice);
// }
// static BOOL  glowReleaseVideoImageNV(GPRELEASEVIDEOIMAGENV fnptr, HPBUFFERARB  hPbuffer, int  iVideoBuffer) {
//   return (*fnptr)(hPbuffer, iVideoBuffer);
// }
// static BOOL  glowResetFrameCountNV(GPRESETFRAMECOUNTNV fnptr, HDC  hDC) {
//   return (*fnptr)(hDC);
// }
// static BOOL  glowRestoreBufferRegionARB(GPRESTOREBUFFERREGIONARB fnptr, HANDLE  hRegion, int  x, int  y, int  width, int  height, int  xSrc, int  ySrc) {
//   return (*fnptr)(hRegion, x, y, width, height, xSrc, ySrc);
// }
// static BOOL  glowSaveBufferRegionARB(GPSAVEBUFFERREGIONARB fnptr, HANDLE  hRegion, int  x, int  y, int  width, int  height) {
//   return (*fnptr)(hRegion, x, y, width, height);
// }
// static BOOL  glowSendPbufferToVideoNV(GPSENDPBUFFERTOVIDEONV fnptr, HPBUFFERARB  hPbuffer, int  iBufferType, unsigned long * pulCounterPbuffer, BOOL  bBlock) {
//   return (*fnptr)(hPbuffer, iBufferType, pulCounterPbuffer, bBlock);
// }
// static BOOL  glowSetDigitalVideoParametersI3D(GPSETDIGITALVIDEOPARAMETERSI3D fnptr, HDC  hDC, int  iAttribute, const int * piValue) {
//   return (*fnptr)(hDC, iAttribute, piValue);
// }
// static BOOL  glowSetGammaTableI3D(GPSETGAMMATABLEI3D fnptr, HDC  hDC, int  iEntries, const USHORT * puRed, const USHORT * puGreen, const USHORT * puBlue) {
//   return (*fnptr)(hDC, iEntries, puRed, puGreen, puBlue);
// }
// static BOOL  glowSetGammaTableParametersI3D(GPSETGAMMATABLEPARAMETERSI3D fnptr, HDC  hDC, int  iAttribute, const int * piValue) {
//   return (*fnptr)(hDC, iAttribute, piValue);
// }
// static int  glowSetLayerPaletteEntries(GPSETLAYERPALETTEENTRIES fnptr, HDC  hdc, int  iLayerPlane, int  iStart, int  cEntries, const COLORREF * pcr) {
//   return (*fnptr)(hdc, iLayerPlane, iStart, cEntries, pcr);
// }
// static BOOL  glowSetPbufferAttribARB(GPSETPBUFFERATTRIBARB fnptr, HPBUFFERARB  hPbuffer, const int * piAttribList) {
//   return (*fnptr)(hPbuffer, piAttribList);
// }
// static BOOL  glowSetStereoEmitterState3DL(GPSETSTEREOEMITTERSTATE3DL fnptr, HDC  hDC, UINT  uState) {
//   return (*fnptr)(hDC, uState);
// }
// static BOOL  glowShareLists(GPSHARELISTS fnptr, HGLRC  hrcSrvShare, HGLRC  hrcSrvSource) {
//   return (*fnptr)(hrcSrvShare, hrcSrvSource);
// }
// static INT64  glowSwapBuffersMscOML(GPSWAPBUFFERSMSCOML fnptr, HDC  hdc, INT64  target_msc, INT64  divisor, INT64  remainder) {
//   return (*fnptr)(hdc, target_msc, divisor, remainder);
// }
// static BOOL  glowSwapIntervalEXT(GPSWAPINTERVALEXT fnptr, int  interval) {
//   return (*fnptr)(interval);
// }
// static BOOL  glowSwapLayerBuffers(GPSWAPLAYERBUFFERS fnptr, HDC  hdc, UINT  fuFlags) {
//   return (*fnptr)(hdc, fuFlags);
// }
// static INT64  glowSwapLayerBuffersMscOML(GPSWAPLAYERBUFFERSMSCOML fnptr, HDC  hdc, int  fuPlanes, INT64  target_msc, INT64  divisor, INT64  remainder) {
//   return (*fnptr)(hdc, fuPlanes, target_msc, divisor, remainder);
// }
// static BOOL  glowUseFontBitmaps(GPUSEFONTBITMAPS fnptr, HDC  hDC, DWORD  first, DWORD  count, DWORD  listBase) {
//   return (*fnptr)(hDC, first, count, listBase);
// }
// static BOOL  glowUseFontBitmapsA(GPUSEFONTBITMAPSA fnptr, HDC  hDC, DWORD  first, DWORD  count, DWORD  listBase) {
//   return (*fnptr)(hDC, first, count, listBase);
// }
// static BOOL  glowUseFontBitmapsW(GPUSEFONTBITMAPSW fnptr, HDC  hDC, DWORD  first, DWORD  count, DWORD  listBase) {
//   return (*fnptr)(hDC, first, count, listBase);
// }
// static BOOL  glowUseFontOutlines(GPUSEFONTOUTLINES fnptr, HDC  hDC, DWORD  first, DWORD  count, DWORD  listBase, FLOAT  deviation, FLOAT  extrusion, int  format, LPGLYPHMETRICSFLOAT  lpgmf) {
//   return (*fnptr)(hDC, first, count, listBase, deviation, extrusion, format, lpgmf);
// }
// static BOOL  glowUseFontOutlinesA(GPUSEFONTOUTLINESA fnptr, HDC  hDC, DWORD  first, DWORD  count, DWORD  listBase, FLOAT  deviation, FLOAT  extrusion, int  format, LPGLYPHMETRICSFLOAT  lpgmf) {
//   return (*fnptr)(hDC, first, count, listBase, deviation, extrusion, format, lpgmf);
// }
// static BOOL  glowUseFontOutlinesW(GPUSEFONTOUTLINESW fnptr, HDC  hDC, DWORD  first, DWORD  count, DWORD  listBase, FLOAT  deviation, FLOAT  extrusion, int  format, LPGLYPHMETRICSFLOAT  lpgmf) {
//   return (*fnptr)(hDC, first, count, listBase, deviation, extrusion, format, lpgmf);
// }
// static BOOL  glowWaitForMscOML(GPWAITFORMSCOML fnptr, HDC  hdc, INT64  target_msc, INT64  divisor, INT64  remainder, INT64 * ust, INT64 * msc, INT64 * sbc) {
//   return (*fnptr)(hdc, target_msc, divisor, remainder, ust, msc, sbc);
// }
// static BOOL  glowWaitForSbcOML(GPWAITFORSBCOML fnptr, HDC  hdc, INT64  target_sbc, INT64 * ust, INT64 * msc, INT64 * sbc) {
//   return (*fnptr)(hdc, target_sbc, ust, msc, sbc);
// }
import "C"
import (
	"errors"
	"unsafe"
)

const (
	ERROR_INCOMPATIBLE_AFFINITY_MASKS_NV       = 0x20D0
	ERROR_INCOMPATIBLE_DEVICE_CONTEXTS_ARB     = 0x2054
	ERROR_INVALID_PIXEL_TYPE_ARB               = 0x2043
	ERROR_INVALID_PIXEL_TYPE_EXT               = 0x2043
	ERROR_INVALID_PROFILE_ARB                  = 0x2096
	ERROR_INVALID_VERSION_ARB                  = 0x2095
	ERROR_MISSING_AFFINITY_MASK_NV             = 0x20D1
	ACCELERATION_ARB                           = 0x2003
	ACCELERATION_EXT                           = 0x2003
	ACCESS_READ_ONLY_NV                        = 0x00000000
	ACCESS_READ_WRITE_NV                       = 0x00000001
	ACCESS_WRITE_DISCARD_NV                    = 0x00000002
	ACCUM_ALPHA_BITS_ARB                       = 0x2021
	ACCUM_ALPHA_BITS_EXT                       = 0x2021
	ACCUM_BITS_ARB                             = 0x201D
	ACCUM_BITS_EXT                             = 0x201D
	ACCUM_BLUE_BITS_ARB                        = 0x2020
	ACCUM_BLUE_BITS_EXT                        = 0x2020
	ACCUM_GREEN_BITS_ARB                       = 0x201F
	ACCUM_GREEN_BITS_EXT                       = 0x201F
	ACCUM_RED_BITS_ARB                         = 0x201E
	ACCUM_RED_BITS_EXT                         = 0x201E
	ALPHA_BITS_ARB                             = 0x201B
	ALPHA_BITS_EXT                             = 0x201B
	ALPHA_SHIFT_ARB                            = 0x201C
	ALPHA_SHIFT_EXT                            = 0x201C
	AUX0_ARB                                   = 0x2087
	AUX1_ARB                                   = 0x2088
	AUX2_ARB                                   = 0x2089
	AUX3_ARB                                   = 0x208A
	AUX4_ARB                                   = 0x208B
	AUX5_ARB                                   = 0x208C
	AUX6_ARB                                   = 0x208D
	AUX7_ARB                                   = 0x208E
	AUX8_ARB                                   = 0x208F
	AUX9_ARB                                   = 0x2090
	AUX_BUFFERS_ARB                            = 0x2024
	AUX_BUFFERS_EXT                            = 0x2024
	BACK_COLOR_BUFFER_BIT_ARB                  = 0x00000002
	BACK_LEFT_ARB                              = 0x2085
	BACK_RIGHT_ARB                             = 0x2086
	BIND_TO_TEXTURE_DEPTH_NV                   = 0x20A3
	BIND_TO_TEXTURE_RECTANGLE_DEPTH_NV         = 0x20A4
	BIND_TO_TEXTURE_RECTANGLE_FLOAT_RGBA_NV    = 0x20B4
	BIND_TO_TEXTURE_RECTANGLE_FLOAT_RGB_NV     = 0x20B3
	BIND_TO_TEXTURE_RECTANGLE_FLOAT_RG_NV      = 0x20B2
	BIND_TO_TEXTURE_RECTANGLE_FLOAT_R_NV       = 0x20B1
	BIND_TO_TEXTURE_RECTANGLE_RGBA_NV          = 0x20A1
	BIND_TO_TEXTURE_RECTANGLE_RGB_NV           = 0x20A0
	BIND_TO_TEXTURE_RGBA_ARB                   = 0x2071
	BIND_TO_TEXTURE_RGB_ARB                    = 0x2070
	BIND_TO_VIDEO_RGBA_NV                      = 0x20C1
	BIND_TO_VIDEO_RGB_AND_DEPTH_NV             = 0x20C2
	BIND_TO_VIDEO_RGB_NV                       = 0x20C0
	BLUE_BITS_ARB                              = 0x2019
	BLUE_BITS_EXT                              = 0x2019
	BLUE_SHIFT_ARB                             = 0x201A
	BLUE_SHIFT_EXT                             = 0x201A
	COLORSPACE_EXT                             = 0x3087
	COLORSPACE_LINEAR_EXT                      = 0x308A
	COLORSPACE_SRGB_EXT                        = 0x3089
	COLOR_BITS_ARB                             = 0x2014
	COLOR_BITS_EXT                             = 0x2014
	COLOR_SAMPLES_NV                           = 0x20B9
	CONTEXT_COMPATIBILITY_PROFILE_BIT_ARB      = 0x00000002
	CONTEXT_CORE_PROFILE_BIT_ARB               = 0x00000001
	CONTEXT_DEBUG_BIT_ARB                      = 0x00000001
	CONTEXT_ES2_PROFILE_BIT_EXT                = 0x00000004
	CONTEXT_ES_PROFILE_BIT_EXT                 = 0x00000004
	CONTEXT_FLAGS_ARB                          = 0x2094
	CONTEXT_FORWARD_COMPATIBLE_BIT_ARB         = 0x00000002
	CONTEXT_LAYER_PLANE_ARB                    = 0x2093
	CONTEXT_MAJOR_VERSION_ARB                  = 0x2091
	CONTEXT_MINOR_VERSION_ARB                  = 0x2092
	CONTEXT_OPENGL_NO_ERROR_ARB                = 0x31B3
	CONTEXT_PROFILE_MASK_ARB                   = 0x9126
	CONTEXT_RELEASE_BEHAVIOR_ARB               = 0x2097
	CONTEXT_RELEASE_BEHAVIOR_FLUSH_ARB         = 0x2098
	CONTEXT_RELEASE_BEHAVIOR_NONE_ARB          = 0
	CONTEXT_RESET_ISOLATION_BIT_ARB            = 0x00000008
	CONTEXT_RESET_NOTIFICATION_STRATEGY_ARB    = 0x8256
	CONTEXT_ROBUST_ACCESS_BIT_ARB              = 0x00000004
	COVERAGE_SAMPLES_NV                        = 0x2042
	CUBE_MAP_FACE_ARB                          = 0x207C
	DEPTH_BITS_ARB                             = 0x2022
	DEPTH_BITS_EXT                             = 0x2022
	DEPTH_BUFFER_BIT_ARB                       = 0x00000004
	DEPTH_COMPONENT_NV                         = 0x20A7
	DEPTH_FLOAT_EXT                            = 0x2040
	DEPTH_TEXTURE_FORMAT_NV                    = 0x20A5
	DIGITAL_VIDEO_CURSOR_ALPHA_FRAMEBUFFER_I3D = 0x2050
	DIGITAL_VIDEO_CURSOR_ALPHA_VALUE_I3D       = 0x2051
	DIGITAL_VIDEO_CURSOR_INCLUDED_I3D          = 0x2052
	DIGITAL_VIDEO_GAMMA_CORRECTED_I3D          = 0x2053
	DOUBLE_BUFFER_ARB                          = 0x2011
	DOUBLE_BUFFER_EXT                          = 0x2011
	DRAW_TO_BITMAP_ARB                         = 0x2002
	DRAW_TO_BITMAP_EXT                         = 0x2002
	DRAW_TO_PBUFFER_ARB                        = 0x202D
	DRAW_TO_PBUFFER_EXT                        = 0x202D
	DRAW_TO_WINDOW_ARB                         = 0x2001
	DRAW_TO_WINDOW_EXT                         = 0x2001
	FLOAT_COMPONENTS_NV                        = 0x20B0
	FONT_LINES                                 = 0
	FONT_POLYGONS                              = 1
	FRAMEBUFFER_SRGB_CAPABLE_ARB               = 0x20A9
	FRAMEBUFFER_SRGB_CAPABLE_EXT               = 0x20A9
	FRONT_COLOR_BUFFER_BIT_ARB                 = 0x00000001
	FRONT_LEFT_ARB                             = 0x2083
	FRONT_RIGHT_ARB                            = 0x2084
	FULL_ACCELERATION_ARB                      = 0x2027
	FULL_ACCELERATION_EXT                      = 0x2027
	GAMMA_EXCLUDE_DESKTOP_I3D                  = 0x204F
	GAMMA_TABLE_SIZE_I3D                       = 0x204E
	GENERIC_ACCELERATION_ARB                   = 0x2026
	GENERIC_ACCELERATION_EXT                   = 0x2026
	GENLOCK_SOURCE_DIGITAL_FIELD_I3D           = 0x2049
	GENLOCK_SOURCE_DIGITAL_SYNC_I3D            = 0x2048
	GENLOCK_SOURCE_EDGE_BOTH_I3D               = 0x204C
	GENLOCK_SOURCE_EDGE_FALLING_I3D            = 0x204A
	GENLOCK_SOURCE_EDGE_RISING_I3D             = 0x204B
	GENLOCK_SOURCE_EXTERNAL_FIELD_I3D          = 0x2046
	GENLOCK_SOURCE_EXTERNAL_SYNC_I3D           = 0x2045
	GENLOCK_SOURCE_EXTERNAL_TTL_I3D            = 0x2047
	GENLOCK_SOURCE_MULTIVIEW_I3D               = 0x2044
	GPU_CLOCK_AMD                              = 0x21A4
	GPU_FASTEST_TARGET_GPUS_AMD                = 0x21A2
	GPU_NUM_PIPES_AMD                          = 0x21A5
	GPU_NUM_RB_AMD                             = 0x21A7
	GPU_NUM_SIMD_AMD                           = 0x21A6
	GPU_NUM_SPI_AMD                            = 0x21A8
	GPU_OPENGL_VERSION_STRING_AMD              = 0x1F02
	GPU_RAM_AMD                                = 0x21A3
	GPU_RENDERER_STRING_AMD                    = 0x1F01
	GPU_VENDOR_AMD                             = 0x1F00
	GREEN_BITS_ARB                             = 0x2017
	GREEN_BITS_EXT                             = 0x2017
	GREEN_SHIFT_ARB                            = 0x2018
	GREEN_SHIFT_EXT                            = 0x2018
	IMAGE_BUFFER_LOCK_I3D                      = 0x00000002
	IMAGE_BUFFER_MIN_ACCESS_I3D                = 0x00000001
	LOSE_CONTEXT_ON_RESET_ARB                  = 0x8252
	MAX_PBUFFER_HEIGHT_ARB                     = 0x2030
	MAX_PBUFFER_HEIGHT_EXT                     = 0x2030
	MAX_PBUFFER_PIXELS_ARB                     = 0x202E
	MAX_PBUFFER_PIXELS_EXT                     = 0x202E
	MAX_PBUFFER_WIDTH_ARB                      = 0x202F
	MAX_PBUFFER_WIDTH_EXT                      = 0x202F
	MIPMAP_LEVEL_ARB                           = 0x207B
	MIPMAP_TEXTURE_ARB                         = 0x2074
	NEED_PALETTE_ARB                           = 0x2004
	NEED_PALETTE_EXT                           = 0x2004
	NEED_SYSTEM_PALETTE_ARB                    = 0x2005
	NEED_SYSTEM_PALETTE_EXT                    = 0x2005
	NO_ACCELERATION_ARB                        = 0x2025
	NO_ACCELERATION_EXT                        = 0x2025
	NO_RESET_NOTIFICATION_ARB                  = 0x8261
	NO_TEXTURE_ARB                             = 0x2077
	NUMBER_OVERLAYS_ARB                        = 0x2008
	NUMBER_OVERLAYS_EXT                        = 0x2008
	NUMBER_PIXEL_FORMATS_ARB                   = 0x2000
	NUMBER_PIXEL_FORMATS_EXT                   = 0x2000
	NUMBER_UNDERLAYS_ARB                       = 0x2009
	NUMBER_UNDERLAYS_EXT                       = 0x2009
	NUM_VIDEO_CAPTURE_SLOTS_NV                 = 0x20CF
	NUM_VIDEO_SLOTS_NV                         = 0x20F0
	OPTIMAL_PBUFFER_HEIGHT_EXT                 = 0x2032
	OPTIMAL_PBUFFER_WIDTH_EXT                  = 0x2031
	PBUFFER_HEIGHT_ARB                         = 0x2035
	PBUFFER_HEIGHT_EXT                         = 0x2035
	PBUFFER_LARGEST_ARB                        = 0x2033
	PBUFFER_LARGEST_EXT                        = 0x2033
	PBUFFER_LOST_ARB                           = 0x2036
	PBUFFER_WIDTH_ARB                          = 0x2034
	PBUFFER_WIDTH_EXT                          = 0x2034
	PIXEL_TYPE_ARB                             = 0x2013
	PIXEL_TYPE_EXT                             = 0x2013
	RED_BITS_ARB                               = 0x2015
	RED_BITS_EXT                               = 0x2015
	RED_SHIFT_ARB                              = 0x2016
	RED_SHIFT_EXT                              = 0x2016
	SAMPLES_3DFX                               = 0x2061
	SAMPLES_ARB                                = 0x2042
	SAMPLES_EXT                                = 0x2042
	SAMPLE_BUFFERS_3DFX                        = 0x2060
	SAMPLE_BUFFERS_ARB                         = 0x2041
	SAMPLE_BUFFERS_EXT                         = 0x2041
	SHARE_ACCUM_ARB                            = 0x200E
	SHARE_ACCUM_EXT                            = 0x200E
	SHARE_DEPTH_ARB                            = 0x200C
	SHARE_DEPTH_EXT                            = 0x200C
	SHARE_STENCIL_ARB                          = 0x200D
	SHARE_STENCIL_EXT                          = 0x200D
	STENCIL_BITS_ARB                           = 0x2023
	STENCIL_BITS_EXT                           = 0x2023
	STENCIL_BUFFER_BIT_ARB                     = 0x00000008
	STEREO_ARB                                 = 0x2012
	STEREO_EMITTER_DISABLE_3DL                 = 0x2056
	STEREO_EMITTER_ENABLE_3DL                  = 0x2055
	STEREO_EXT                                 = 0x2012
	STEREO_POLARITY_INVERT_3DL                 = 0x2058
	STEREO_POLARITY_NORMAL_3DL                 = 0x2057
	SUPPORT_GDI_ARB                            = 0x200F
	SUPPORT_GDI_EXT                            = 0x200F
	SUPPORT_OPENGL_ARB                         = 0x2010
	SUPPORT_OPENGL_EXT                         = 0x2010
	SWAP_COPY_ARB                              = 0x2029
	SWAP_COPY_EXT                              = 0x2029
	SWAP_EXCHANGE_ARB                          = 0x2028
	SWAP_EXCHANGE_EXT                          = 0x2028
	SWAP_LAYER_BUFFERS_ARB                     = 0x2006
	SWAP_LAYER_BUFFERS_EXT                     = 0x2006
	SWAP_MAIN_PLANE                            = 0x00000001
	SWAP_METHOD_ARB                            = 0x2007
	SWAP_METHOD_EXT                            = 0x2007
	SWAP_OVERLAY1                              = 0x00000002
	SWAP_OVERLAY10                             = 0x00000400
	SWAP_OVERLAY11                             = 0x00000800
	SWAP_OVERLAY12                             = 0x00001000
	SWAP_OVERLAY13                             = 0x00002000
	SWAP_OVERLAY14                             = 0x00004000
	SWAP_OVERLAY15                             = 0x00008000
	SWAP_OVERLAY2                              = 0x00000004
	SWAP_OVERLAY3                              = 0x00000008
	SWAP_OVERLAY4                              = 0x00000010
	SWAP_OVERLAY5                              = 0x00000020
	SWAP_OVERLAY6                              = 0x00000040
	SWAP_OVERLAY7                              = 0x00000080
	SWAP_OVERLAY8                              = 0x00000100
	SWAP_OVERLAY9                              = 0x00000200
	SWAP_UNDEFINED_ARB                         = 0x202A
	SWAP_UNDEFINED_EXT                         = 0x202A
	SWAP_UNDERLAY1                             = 0x00010000
	SWAP_UNDERLAY10                            = 0x02000000
	SWAP_UNDERLAY11                            = 0x04000000
	SWAP_UNDERLAY12                            = 0x08000000
	SWAP_UNDERLAY13                            = 0x10000000
	SWAP_UNDERLAY14                            = 0x20000000
	SWAP_UNDERLAY15                            = 0x40000000
	SWAP_UNDERLAY2                             = 0x00020000
	SWAP_UNDERLAY3                             = 0x00040000
	SWAP_UNDERLAY4                             = 0x00080000
	SWAP_UNDERLAY5                             = 0x00100000
	SWAP_UNDERLAY6                             = 0x00200000
	SWAP_UNDERLAY7                             = 0x00400000
	SWAP_UNDERLAY8                             = 0x00800000
	SWAP_UNDERLAY9                             = 0x01000000
	TEXTURE_1D_ARB                             = 0x2079
	TEXTURE_2D_ARB                             = 0x207A
	TEXTURE_CUBE_MAP_ARB                       = 0x2078
	TEXTURE_CUBE_MAP_NEGATIVE_X_ARB            = 0x207E
	TEXTURE_CUBE_MAP_NEGATIVE_Y_ARB            = 0x2080
	TEXTURE_CUBE_MAP_NEGATIVE_Z_ARB            = 0x2082
	TEXTURE_CUBE_MAP_POSITIVE_X_ARB            = 0x207D
	TEXTURE_CUBE_MAP_POSITIVE_Y_ARB            = 0x207F
	TEXTURE_CUBE_MAP_POSITIVE_Z_ARB            = 0x2081
	TEXTURE_DEPTH_COMPONENT_NV                 = 0x20A6
	TEXTURE_FLOAT_RGBA_NV                      = 0x20B8
	TEXTURE_FLOAT_RGB_NV                       = 0x20B7
	TEXTURE_FLOAT_RG_NV                        = 0x20B6
	TEXTURE_FLOAT_R_NV                         = 0x20B5
	TEXTURE_FORMAT_ARB                         = 0x2072
	TEXTURE_RECTANGLE_NV                       = 0x20A2
	TEXTURE_RGBA_ARB                           = 0x2076
	TEXTURE_RGB_ARB                            = 0x2075
	TEXTURE_TARGET_ARB                         = 0x2073
	TRANSPARENT_ALPHA_VALUE_ARB                = 0x203A
	TRANSPARENT_ARB                            = 0x200A
	TRANSPARENT_BLUE_VALUE_ARB                 = 0x2039
	TRANSPARENT_EXT                            = 0x200A
	TRANSPARENT_GREEN_VALUE_ARB                = 0x2038
	TRANSPARENT_INDEX_VALUE_ARB                = 0x203B
	TRANSPARENT_RED_VALUE_ARB                  = 0x2037
	TRANSPARENT_VALUE_EXT                      = 0x200B
	TYPE_COLORINDEX_ARB                        = 0x202C
	TYPE_COLORINDEX_EXT                        = 0x202C
	TYPE_RGBA_ARB                              = 0x202B
	TYPE_RGBA_EXT                              = 0x202B
	TYPE_RGBA_FLOAT_ARB                        = 0x21A0
	TYPE_RGBA_FLOAT_ATI                        = 0x21A0
	TYPE_RGBA_UNSIGNED_FLOAT_EXT               = 0x20A8
	UNIQUE_ID_NV                               = 0x20CE
	VIDEO_OUT_ALPHA_NV                         = 0x20C4
	VIDEO_OUT_COLOR_AND_ALPHA_NV               = 0x20C6
	VIDEO_OUT_COLOR_AND_DEPTH_NV               = 0x20C7
	VIDEO_OUT_COLOR_NV                         = 0x20C3
	VIDEO_OUT_DEPTH_NV                         = 0x20C5
	VIDEO_OUT_FIELD_1                          = 0x20C9
	VIDEO_OUT_FIELD_2                          = 0x20CA
	VIDEO_OUT_FRAME                            = 0x20C8
	VIDEO_OUT_STACKED_FIELDS_1_2               = 0x20CB
	VIDEO_OUT_STACKED_FIELDS_2_1               = 0x20CC
)

var (
	gpChoosePixelFormat                 C.GPCHOOSEPIXELFORMAT
	gpDescribePixelFormat               C.GPDESCRIBEPIXELFORMAT
	gpGetEnhMetaFilePixelFormat         C.GPGETENHMETAFILEPIXELFORMAT
	gpGetPixelFormat                    C.GPGETPIXELFORMAT
	gpSetPixelFormat                    C.GPSETPIXELFORMAT
	gpSwapBuffers                       C.GPSWAPBUFFERS
	gpAllocateMemoryNV                  C.GPALLOCATEMEMORYNV
	gpAssociateImageBufferEventsI3D     C.GPASSOCIATEIMAGEBUFFEREVENTSI3D
	gpBeginFrameTrackingI3D             C.GPBEGINFRAMETRACKINGI3D
	gpBindDisplayColorTableEXT          C.GPBINDDISPLAYCOLORTABLEEXT
	gpBindSwapBarrierNV                 C.GPBINDSWAPBARRIERNV
	gpBindTexImageARB                   C.GPBINDTEXIMAGEARB
	gpBindVideoCaptureDeviceNV          C.GPBINDVIDEOCAPTUREDEVICENV
	gpBindVideoDeviceNV                 C.GPBINDVIDEODEVICENV
	gpBindVideoImageNV                  C.GPBINDVIDEOIMAGENV
	gpBlitContextFramebufferAMD         C.GPBLITCONTEXTFRAMEBUFFERAMD
	gpChoosePixelFormatARB              C.GPCHOOSEPIXELFORMATARB
	gpChoosePixelFormatEXT              C.GPCHOOSEPIXELFORMATEXT
	gpCopyContext                       C.GPCOPYCONTEXT
	gpCopyImageSubDataNV                C.GPCOPYIMAGESUBDATANV
	gpCreateAffinityDCNV                C.GPCREATEAFFINITYDCNV
	gpCreateAssociatedContextAMD        C.GPCREATEASSOCIATEDCONTEXTAMD
	gpCreateAssociatedContextAttribsAMD C.GPCREATEASSOCIATEDCONTEXTATTRIBSAMD
	gpCreateBufferRegionARB             C.GPCREATEBUFFERREGIONARB
	gpCreateContext                     C.GPCREATECONTEXT
	gpCreateContextAttribsARB           C.GPCREATECONTEXTATTRIBSARB
	gpCreateDisplayColorTableEXT        C.GPCREATEDISPLAYCOLORTABLEEXT
	gpCreateImageBufferI3D              C.GPCREATEIMAGEBUFFERI3D
	gpCreateLayerContext                C.GPCREATELAYERCONTEXT
	gpCreatePbufferARB                  C.GPCREATEPBUFFERARB
	gpCreatePbufferEXT                  C.GPCREATEPBUFFEREXT
	gpDXCloseDeviceNV                   C.GPDXCLOSEDEVICENV
	gpDXLockObjectsNV                   C.GPDXLOCKOBJECTSNV
	gpDXObjectAccessNV                  C.GPDXOBJECTACCESSNV
	gpDXOpenDeviceNV                    C.GPDXOPENDEVICENV
	gpDXRegisterObjectNV                C.GPDXREGISTEROBJECTNV
	gpDXSetResourceShareHandleNV        C.GPDXSETRESOURCESHAREHANDLENV
	gpDXUnlockObjectsNV                 C.GPDXUNLOCKOBJECTSNV
	gpDXUnregisterObjectNV              C.GPDXUNREGISTEROBJECTNV
	gpDelayBeforeSwapNV                 C.GPDELAYBEFORESWAPNV
	gpDeleteAssociatedContextAMD        C.GPDELETEASSOCIATEDCONTEXTAMD
	gpDeleteBufferRegionARB             C.GPDELETEBUFFERREGIONARB
	gpDeleteContext                     C.GPDELETECONTEXT
	gpDeleteDCNV                        C.GPDELETEDCNV
	gpDescribeLayerPlane                C.GPDESCRIBELAYERPLANE
	gpDestroyDisplayColorTableEXT       C.GPDESTROYDISPLAYCOLORTABLEEXT
	gpDestroyImageBufferI3D             C.GPDESTROYIMAGEBUFFERI3D
	gpDestroyPbufferARB                 C.GPDESTROYPBUFFERARB
	gpDestroyPbufferEXT                 C.GPDESTROYPBUFFEREXT
	gpDisableFrameLockI3D               C.GPDISABLEFRAMELOCKI3D
	gpDisableGenlockI3D                 C.GPDISABLEGENLOCKI3D
	gpEnableFrameLockI3D                C.GPENABLEFRAMELOCKI3D
	gpEnableGenlockI3D                  C.GPENABLEGENLOCKI3D
	gpEndFrameTrackingI3D               C.GPENDFRAMETRACKINGI3D
	gpEnumGpuDevicesNV                  C.GPENUMGPUDEVICESNV
	gpEnumGpusFromAffinityDCNV          C.GPENUMGPUSFROMAFFINITYDCNV
	gpEnumGpusNV                        C.GPENUMGPUSNV
	gpEnumerateVideoCaptureDevicesNV    C.GPENUMERATEVIDEOCAPTUREDEVICESNV
	gpEnumerateVideoDevicesNV           C.GPENUMERATEVIDEODEVICESNV
	gpFreeMemoryNV                      C.GPFREEMEMORYNV
	gpGenlockSampleRateI3D              C.GPGENLOCKSAMPLERATEI3D
	gpGenlockSourceDelayI3D             C.GPGENLOCKSOURCEDELAYI3D
	gpGenlockSourceEdgeI3D              C.GPGENLOCKSOURCEEDGEI3D
	gpGenlockSourceI3D                  C.GPGENLOCKSOURCEI3D
	gpGetContextGPUIDAMD                C.GPGETCONTEXTGPUIDAMD
	gpGetCurrentAssociatedContextAMD    C.GPGETCURRENTASSOCIATEDCONTEXTAMD
	gpGetCurrentContext                 C.GPGETCURRENTCONTEXT
	gpGetCurrentDC                      C.GPGETCURRENTDC
	gpGetCurrentReadDCARB               C.GPGETCURRENTREADDCARB
	gpGetCurrentReadDCEXT               C.GPGETCURRENTREADDCEXT
	gpGetDigitalVideoParametersI3D      C.GPGETDIGITALVIDEOPARAMETERSI3D
	gpGetExtensionsStringARB            C.GPGETEXTENSIONSSTRINGARB
	gpGetExtensionsStringEXT            C.GPGETEXTENSIONSSTRINGEXT
	gpGetFrameUsageI3D                  C.GPGETFRAMEUSAGEI3D
	gpGetGPUIDsAMD                      C.GPGETGPUIDSAMD
	gpGetGPUInfoAMD                     C.GPGETGPUINFOAMD
	gpGetGammaTableI3D                  C.GPGETGAMMATABLEI3D
	gpGetGammaTableParametersI3D        C.GPGETGAMMATABLEPARAMETERSI3D
	gpGetGenlockSampleRateI3D           C.GPGETGENLOCKSAMPLERATEI3D
	gpGetGenlockSourceDelayI3D          C.GPGETGENLOCKSOURCEDELAYI3D
	gpGetGenlockSourceEdgeI3D           C.GPGETGENLOCKSOURCEEDGEI3D
	gpGetGenlockSourceI3D               C.GPGETGENLOCKSOURCEI3D
	gpGetLayerPaletteEntries            C.GPGETLAYERPALETTEENTRIES
	gpGetMscRateOML                     C.GPGETMSCRATEOML
	gpGetPbufferDCARB                   C.GPGETPBUFFERDCARB
	gpGetPbufferDCEXT                   C.GPGETPBUFFERDCEXT
	gpGetPixelFormatAttribfvARB         C.GPGETPIXELFORMATATTRIBFVARB
	gpGetPixelFormatAttribfvEXT         C.GPGETPIXELFORMATATTRIBFVEXT
	gpGetPixelFormatAttribivARB         C.GPGETPIXELFORMATATTRIBIVARB
	gpGetPixelFormatAttribivEXT         C.GPGETPIXELFORMATATTRIBIVEXT
	gpGetProcAddress                    C.GPGETPROCADDRESS
	gpGetSwapIntervalEXT                C.GPGETSWAPINTERVALEXT
	gpGetSyncValuesOML                  C.GPGETSYNCVALUESOML
	gpGetVideoDeviceNV                  C.GPGETVIDEODEVICENV
	gpGetVideoInfoNV                    C.GPGETVIDEOINFONV
	gpIsEnabledFrameLockI3D             C.GPISENABLEDFRAMELOCKI3D
	gpIsEnabledGenlockI3D               C.GPISENABLEDGENLOCKI3D
	gpJoinSwapGroupNV                   C.GPJOINSWAPGROUPNV
	gpLoadDisplayColorTableEXT          C.GPLOADDISPLAYCOLORTABLEEXT
	gpLockVideoCaptureDeviceNV          C.GPLOCKVIDEOCAPTUREDEVICENV
	gpMakeAssociatedContextCurrentAMD   C.GPMAKEASSOCIATEDCONTEXTCURRENTAMD
	gpMakeContextCurrentARB             C.GPMAKECONTEXTCURRENTARB
	gpMakeContextCurrentEXT             C.GPMAKECONTEXTCURRENTEXT
	gpMakeCurrent                       C.GPMAKECURRENT
	gpQueryCurrentContextNV             C.GPQUERYCURRENTCONTEXTNV
	gpQueryFrameCountNV                 C.GPQUERYFRAMECOUNTNV
	gpQueryFrameLockMasterI3D           C.GPQUERYFRAMELOCKMASTERI3D
	gpQueryFrameTrackingI3D             C.GPQUERYFRAMETRACKINGI3D
	gpQueryGenlockMaxSourceDelayI3D     C.GPQUERYGENLOCKMAXSOURCEDELAYI3D
	gpQueryMaxSwapGroupsNV              C.GPQUERYMAXSWAPGROUPSNV
	gpQueryPbufferARB                   C.GPQUERYPBUFFERARB
	gpQueryPbufferEXT                   C.GPQUERYPBUFFEREXT
	gpQuerySwapGroupNV                  C.GPQUERYSWAPGROUPNV
	gpQueryVideoCaptureDeviceNV         C.GPQUERYVIDEOCAPTUREDEVICENV
	gpRealizeLayerPalette               C.GPREALIZELAYERPALETTE
	gpReleaseImageBufferEventsI3D       C.GPRELEASEIMAGEBUFFEREVENTSI3D
	gpReleasePbufferDCARB               C.GPRELEASEPBUFFERDCARB
	gpReleasePbufferDCEXT               C.GPRELEASEPBUFFERDCEXT
	gpReleaseTexImageARB                C.GPRELEASETEXIMAGEARB
	gpReleaseVideoCaptureDeviceNV       C.GPRELEASEVIDEOCAPTUREDEVICENV
	gpReleaseVideoDeviceNV              C.GPRELEASEVIDEODEVICENV
	gpReleaseVideoImageNV               C.GPRELEASEVIDEOIMAGENV
	gpResetFrameCountNV                 C.GPRESETFRAMECOUNTNV
	gpRestoreBufferRegionARB            C.GPRESTOREBUFFERREGIONARB
	gpSaveBufferRegionARB               C.GPSAVEBUFFERREGIONARB
	gpSendPbufferToVideoNV              C.GPSENDPBUFFERTOVIDEONV
	gpSetDigitalVideoParametersI3D      C.GPSETDIGITALVIDEOPARAMETERSI3D
	gpSetGammaTableI3D                  C.GPSETGAMMATABLEI3D
	gpSetGammaTableParametersI3D        C.GPSETGAMMATABLEPARAMETERSI3D
	gpSetLayerPaletteEntries            C.GPSETLAYERPALETTEENTRIES
	gpSetPbufferAttribARB               C.GPSETPBUFFERATTRIBARB
	gpSetStereoEmitterState3DL          C.GPSETSTEREOEMITTERSTATE3DL
	gpShareLists                        C.GPSHARELISTS
	gpSwapBuffersMscOML                 C.GPSWAPBUFFERSMSCOML
	gpSwapIntervalEXT                   C.GPSWAPINTERVALEXT
	gpSwapLayerBuffers                  C.GPSWAPLAYERBUFFERS
	gpSwapLayerBuffersMscOML            C.GPSWAPLAYERBUFFERSMSCOML
	gpUseFontBitmaps                    C.GPUSEFONTBITMAPS
	gpUseFontBitmapsA                   C.GPUSEFONTBITMAPSA
	gpUseFontBitmapsW                   C.GPUSEFONTBITMAPSW
	gpUseFontOutlines                   C.GPUSEFONTOUTLINES
	gpUseFontOutlinesA                  C.GPUSEFONTOUTLINESA
	gpUseFontOutlinesW                  C.GPUSEFONTOUTLINESW
	gpWaitForMscOML                     C.GPWAITFORMSCOML
	gpWaitForSbcOML                     C.GPWAITFORSBCOML
)

// Helper functions
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
func ChoosePixelFormat(hDc C.HDC, pPfd *C.PIXELFORMATDESCRIPTOR) C.int {
	ret := C.glowChoosePixelFormat(gpChoosePixelFormat, (C.HDC)(hDc), (*C.PIXELFORMATDESCRIPTOR)(unsafe.Pointer(pPfd)))
	return (C.int)(ret)
}
func DescribePixelFormat(hdc C.HDC, ipfd C.int, cjpfd C.UINT, ppfd *C.PIXELFORMATDESCRIPTOR) C.int {
	ret := C.glowDescribePixelFormat(gpDescribePixelFormat, (C.HDC)(hdc), (C.int)(ipfd), (C.UINT)(cjpfd), (*C.PIXELFORMATDESCRIPTOR)(unsafe.Pointer(ppfd)))
	return (C.int)(ret)
}
func GetEnhMetaFilePixelFormat(hemf C.HENHMETAFILE, ppfd *C.PIXELFORMATDESCRIPTOR) C.UINT {
	ret := C.glowGetEnhMetaFilePixelFormat(gpGetEnhMetaFilePixelFormat, (C.HENHMETAFILE)(hemf), (*C.PIXELFORMATDESCRIPTOR)(unsafe.Pointer(ppfd)))
	return (C.UINT)(ret)
}
func GetPixelFormat(hdc C.HDC) C.int {
	ret := C.glowGetPixelFormat(gpGetPixelFormat, (C.HDC)(hdc))
	return (C.int)(ret)
}
func SetPixelFormat(hdc C.HDC, ipfd C.int, ppfd *C.PIXELFORMATDESCRIPTOR) C.BOOL {
	ret := C.glowSetPixelFormat(gpSetPixelFormat, (C.HDC)(hdc), (C.int)(ipfd), (*C.PIXELFORMATDESCRIPTOR)(unsafe.Pointer(ppfd)))
	return (C.BOOL)(ret)
}
func SwapBuffers(hdc C.HDC) C.BOOL {
	ret := C.glowSwapBuffers(gpSwapBuffers, (C.HDC)(hdc))
	return (C.BOOL)(ret)
}
func AllocateMemoryNV(size int32, readfreq float32, writefreq float32, priority float32) unsafe.Pointer {
	ret := C.glowAllocateMemoryNV(gpAllocateMemoryNV, (C.GLsizei)(size), (C.GLfloat)(readfreq), (C.GLfloat)(writefreq), (C.GLfloat)(priority))
	return (unsafe.Pointer)(ret)
}
func AssociateImageBufferEventsI3D(hDC C.HDC, pEvent *C.HANDLE, pAddress *C.LPVOID, pSize *C.DWORD, count C.UINT) C.BOOL {
	ret := C.glowAssociateImageBufferEventsI3D(gpAssociateImageBufferEventsI3D, (C.HDC)(hDC), (*C.HANDLE)(unsafe.Pointer(pEvent)), (*C.LPVOID)(unsafe.Pointer(pAddress)), (*C.DWORD)(unsafe.Pointer(pSize)), (C.UINT)(count))
	return (C.BOOL)(ret)
}
func BeginFrameTrackingI3D() C.BOOL {
	ret := C.glowBeginFrameTrackingI3D(gpBeginFrameTrackingI3D)
	return (C.BOOL)(ret)
}
func BindDisplayColorTableEXT(id uint16) bool {
	ret := C.glowBindDisplayColorTableEXT(gpBindDisplayColorTableEXT, (C.GLushort)(id))
	return ret == TRUE
}
func BindSwapBarrierNV(group uint32, barrier uint32) C.BOOL {
	ret := C.glowBindSwapBarrierNV(gpBindSwapBarrierNV, (C.GLuint)(group), (C.GLuint)(barrier))
	return (C.BOOL)(ret)
}
func BindTexImageARB(hPbuffer C.HPBUFFERARB, iBuffer C.int) C.BOOL {
	ret := C.glowBindTexImageARB(gpBindTexImageARB, (C.HPBUFFERARB)(hPbuffer), (C.int)(iBuffer))
	return (C.BOOL)(ret)
}
func BindVideoCaptureDeviceNV(uVideoSlot C.UINT, hDevice C.HVIDEOINPUTDEVICENV) C.BOOL {
	ret := C.glowBindVideoCaptureDeviceNV(gpBindVideoCaptureDeviceNV, (C.UINT)(uVideoSlot), (C.HVIDEOINPUTDEVICENV)(hDevice))
	return (C.BOOL)(ret)
}
func BindVideoDeviceNV(hDC C.HDC, uVideoSlot C.unsignedint, hVideoDevice C.HVIDEOOUTPUTDEVICENV, piAttribList *C.int) C.BOOL {
	ret := C.glowBindVideoDeviceNV(gpBindVideoDeviceNV, (C.HDC)(hDC), (C.unsignedint)(uVideoSlot), (C.HVIDEOOUTPUTDEVICENV)(hVideoDevice), (*C.int)(unsafe.Pointer(piAttribList)))
	return (C.BOOL)(ret)
}
func BindVideoImageNV(hVideoDevice C.HPVIDEODEV, hPbuffer C.HPBUFFERARB, iVideoBuffer C.int) C.BOOL {
	ret := C.glowBindVideoImageNV(gpBindVideoImageNV, (C.HPVIDEODEV)(hVideoDevice), (C.HPBUFFERARB)(hPbuffer), (C.int)(iVideoBuffer))
	return (C.BOOL)(ret)
}
func BlitContextFramebufferAMD(dstCtx C.HGLRC, srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter uint32) C.VOID {
	ret := C.glowBlitContextFramebufferAMD(gpBlitContextFramebufferAMD, (C.HGLRC)(dstCtx), (C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter))
	return (C.VOID)(ret)
}
func ChoosePixelFormatARB(hdc C.HDC, piAttribIList *C.int, pfAttribFList *C.FLOAT, nMaxFormats C.UINT, piFormats *C.int, nNumFormats *C.UINT) C.BOOL {
	ret := C.glowChoosePixelFormatARB(gpChoosePixelFormatARB, (C.HDC)(hdc), (*C.int)(unsafe.Pointer(piAttribIList)), (*C.FLOAT)(unsafe.Pointer(pfAttribFList)), (C.UINT)(nMaxFormats), (*C.int)(unsafe.Pointer(piFormats)), (*C.UINT)(unsafe.Pointer(nNumFormats)))
	return (C.BOOL)(ret)
}
func ChoosePixelFormatEXT(hdc C.HDC, piAttribIList *C.int, pfAttribFList *C.FLOAT, nMaxFormats C.UINT, piFormats *C.int, nNumFormats *C.UINT) C.BOOL {
	ret := C.glowChoosePixelFormatEXT(gpChoosePixelFormatEXT, (C.HDC)(hdc), (*C.int)(unsafe.Pointer(piAttribIList)), (*C.FLOAT)(unsafe.Pointer(pfAttribFList)), (C.UINT)(nMaxFormats), (*C.int)(unsafe.Pointer(piFormats)), (*C.UINT)(unsafe.Pointer(nNumFormats)))
	return (C.BOOL)(ret)
}
func CopyContext(hglrcSrc C.HGLRC, hglrcDst C.HGLRC, mask C.UINT) C.BOOL {
	ret := C.glowCopyContext(gpCopyContext, (C.HGLRC)(hglrcSrc), (C.HGLRC)(hglrcDst), (C.UINT)(mask))
	return (C.BOOL)(ret)
}
func CopyImageSubDataNV(hSrcRC C.HGLRC, srcName uint32, srcTarget uint32, srcLevel int32, srcX int32, srcY int32, srcZ int32, hDstRC C.HGLRC, dstName uint32, dstTarget uint32, dstLevel int32, dstX int32, dstY int32, dstZ int32, width int32, height int32, depth int32) C.BOOL {
	ret := C.glowCopyImageSubDataNV(gpCopyImageSubDataNV, (C.HGLRC)(hSrcRC), (C.GLuint)(srcName), (C.GLenum)(srcTarget), (C.GLint)(srcLevel), (C.GLint)(srcX), (C.GLint)(srcY), (C.GLint)(srcZ), (C.HGLRC)(hDstRC), (C.GLuint)(dstName), (C.GLenum)(dstTarget), (C.GLint)(dstLevel), (C.GLint)(dstX), (C.GLint)(dstY), (C.GLint)(dstZ), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
	return (C.BOOL)(ret)
}
func CreateAffinityDCNV(phGpuList *C.HGPUNV) C.HDC {
	ret := C.glowCreateAffinityDCNV(gpCreateAffinityDCNV, (*C.HGPUNV)(unsafe.Pointer(phGpuList)))
	return (C.HDC)(ret)
}
func CreateAssociatedContextAMD(id C.UINT) C.HGLRC {
	ret := C.glowCreateAssociatedContextAMD(gpCreateAssociatedContextAMD, (C.UINT)(id))
	return (C.HGLRC)(ret)
}
func CreateAssociatedContextAttribsAMD(id C.UINT, hShareContext C.HGLRC, attribList *C.int) C.HGLRC {
	ret := C.glowCreateAssociatedContextAttribsAMD(gpCreateAssociatedContextAttribsAMD, (C.UINT)(id), (C.HGLRC)(hShareContext), (*C.int)(unsafe.Pointer(attribList)))
	return (C.HGLRC)(ret)
}
func CreateBufferRegionARB(hDC C.HDC, iLayerPlane C.int, uType C.UINT) C.HANDLE {
	ret := C.glowCreateBufferRegionARB(gpCreateBufferRegionARB, (C.HDC)(hDC), (C.int)(iLayerPlane), (C.UINT)(uType))
	return (C.HANDLE)(ret)
}
func CreateContext(hDc C.HDC) C.HGLRC {
	ret := C.glowCreateContext(gpCreateContext, (C.HDC)(hDc))
	return (C.HGLRC)(ret)
}
func CreateContextAttribsARB(hDC C.HDC, hShareContext C.HGLRC, attribList *C.int) C.HGLRC {
	ret := C.glowCreateContextAttribsARB(gpCreateContextAttribsARB, (C.HDC)(hDC), (C.HGLRC)(hShareContext), (*C.int)(unsafe.Pointer(attribList)))
	return (C.HGLRC)(ret)
}
func CreateDisplayColorTableEXT(id uint16) bool {
	ret := C.glowCreateDisplayColorTableEXT(gpCreateDisplayColorTableEXT, (C.GLushort)(id))
	return ret == TRUE
}
func CreateImageBufferI3D(hDC C.HDC, dwSize C.DWORD, uFlags C.UINT) C.LPVOID {
	ret := C.glowCreateImageBufferI3D(gpCreateImageBufferI3D, (C.HDC)(hDC), (C.DWORD)(dwSize), (C.UINT)(uFlags))
	return (C.LPVOID)(ret)
}
func CreateLayerContext(hDc C.HDC, level C.int) C.HGLRC {
	ret := C.glowCreateLayerContext(gpCreateLayerContext, (C.HDC)(hDc), (C.int)(level))
	return (C.HGLRC)(ret)
}
func CreatePbufferARB(hDC C.HDC, iPixelFormat C.int, iWidth C.int, iHeight C.int, piAttribList *C.int) C.HPBUFFERARB {
	ret := C.glowCreatePbufferARB(gpCreatePbufferARB, (C.HDC)(hDC), (C.int)(iPixelFormat), (C.int)(iWidth), (C.int)(iHeight), (*C.int)(unsafe.Pointer(piAttribList)))
	return (C.HPBUFFERARB)(ret)
}
func CreatePbufferEXT(hDC C.HDC, iPixelFormat C.int, iWidth C.int, iHeight C.int, piAttribList *C.int) C.HPBUFFEREXT {
	ret := C.glowCreatePbufferEXT(gpCreatePbufferEXT, (C.HDC)(hDC), (C.int)(iPixelFormat), (C.int)(iWidth), (C.int)(iHeight), (*C.int)(unsafe.Pointer(piAttribList)))
	return (C.HPBUFFEREXT)(ret)
}
func DXCloseDeviceNV(hDevice C.HANDLE) C.BOOL {
	ret := C.glowDXCloseDeviceNV(gpDXCloseDeviceNV, (C.HANDLE)(hDevice))
	return (C.BOOL)(ret)
}
func DXLockObjectsNV(hDevice C.HANDLE, count int32, hObjects *C.HANDLE) C.BOOL {
	ret := C.glowDXLockObjectsNV(gpDXLockObjectsNV, (C.HANDLE)(hDevice), (C.GLint)(count), (*C.HANDLE)(unsafe.Pointer(hObjects)))
	return (C.BOOL)(ret)
}
func DXObjectAccessNV(hObject C.HANDLE, access uint32) C.BOOL {
	ret := C.glowDXObjectAccessNV(gpDXObjectAccessNV, (C.HANDLE)(hObject), (C.GLenum)(access))
	return (C.BOOL)(ret)
}
func DXOpenDeviceNV(dxDevice unsafe.Pointer) C.HANDLE {
	ret := C.glowDXOpenDeviceNV(gpDXOpenDeviceNV, dxDevice)
	return (C.HANDLE)(ret)
}
func DXRegisterObjectNV(hDevice C.HANDLE, dxObject unsafe.Pointer, name uint32, xtype uint32, access uint32) C.HANDLE {
	ret := C.glowDXRegisterObjectNV(gpDXRegisterObjectNV, (C.HANDLE)(hDevice), dxObject, (C.GLuint)(name), (C.GLenum)(xtype), (C.GLenum)(access))
	return (C.HANDLE)(ret)
}
func DXSetResourceShareHandleNV(dxObject unsafe.Pointer, shareHandle C.HANDLE) C.BOOL {
	ret := C.glowDXSetResourceShareHandleNV(gpDXSetResourceShareHandleNV, dxObject, (C.HANDLE)(shareHandle))
	return (C.BOOL)(ret)
}
func DXUnlockObjectsNV(hDevice C.HANDLE, count int32, hObjects *C.HANDLE) C.BOOL {
	ret := C.glowDXUnlockObjectsNV(gpDXUnlockObjectsNV, (C.HANDLE)(hDevice), (C.GLint)(count), (*C.HANDLE)(unsafe.Pointer(hObjects)))
	return (C.BOOL)(ret)
}
func DXUnregisterObjectNV(hDevice C.HANDLE, hObject C.HANDLE) C.BOOL {
	ret := C.glowDXUnregisterObjectNV(gpDXUnregisterObjectNV, (C.HANDLE)(hDevice), (C.HANDLE)(hObject))
	return (C.BOOL)(ret)
}
func DelayBeforeSwapNV(hDC C.HDC, seconds float32) C.BOOL {
	ret := C.glowDelayBeforeSwapNV(gpDelayBeforeSwapNV, (C.HDC)(hDC), (C.GLfloat)(seconds))
	return (C.BOOL)(ret)
}
func DeleteAssociatedContextAMD(hglrc C.HGLRC) C.BOOL {
	ret := C.glowDeleteAssociatedContextAMD(gpDeleteAssociatedContextAMD, (C.HGLRC)(hglrc))
	return (C.BOOL)(ret)
}
func DeleteBufferRegionARB(hRegion C.HANDLE) C.VOID {
	ret := C.glowDeleteBufferRegionARB(gpDeleteBufferRegionARB, (C.HANDLE)(hRegion))
	return (C.VOID)(ret)
}
func DeleteContext(oldContext C.HGLRC) C.BOOL {
	ret := C.glowDeleteContext(gpDeleteContext, (C.HGLRC)(oldContext))
	return (C.BOOL)(ret)
}
func DeleteDCNV(hdc C.HDC) C.BOOL {
	ret := C.glowDeleteDCNV(gpDeleteDCNV, (C.HDC)(hdc))
	return (C.BOOL)(ret)
}
func DescribeLayerPlane(hDc C.HDC, pixelFormat C.int, layerPlane C.int, nBytes C.UINT, plpd *C.LAYERPLANEDESCRIPTOR) C.BOOL {
	ret := C.glowDescribeLayerPlane(gpDescribeLayerPlane, (C.HDC)(hDc), (C.int)(pixelFormat), (C.int)(layerPlane), (C.UINT)(nBytes), (*C.LAYERPLANEDESCRIPTOR)(unsafe.Pointer(plpd)))
	return (C.BOOL)(ret)
}
func DestroyDisplayColorTableEXT(id uint16) C.VOID {
	ret := C.glowDestroyDisplayColorTableEXT(gpDestroyDisplayColorTableEXT, (C.GLushort)(id))
	return (C.VOID)(ret)
}
func DestroyImageBufferI3D(hDC C.HDC, pAddress C.LPVOID) C.BOOL {
	ret := C.glowDestroyImageBufferI3D(gpDestroyImageBufferI3D, (C.HDC)(hDC), (C.LPVOID)(pAddress))
	return (C.BOOL)(ret)
}
func DestroyPbufferARB(hPbuffer C.HPBUFFERARB) C.BOOL {
	ret := C.glowDestroyPbufferARB(gpDestroyPbufferARB, (C.HPBUFFERARB)(hPbuffer))
	return (C.BOOL)(ret)
}
func DestroyPbufferEXT(hPbuffer C.HPBUFFEREXT) C.BOOL {
	ret := C.glowDestroyPbufferEXT(gpDestroyPbufferEXT, (C.HPBUFFEREXT)(hPbuffer))
	return (C.BOOL)(ret)
}
func DisableFrameLockI3D() C.BOOL {
	ret := C.glowDisableFrameLockI3D(gpDisableFrameLockI3D)
	return (C.BOOL)(ret)
}
func DisableGenlockI3D(hDC C.HDC) C.BOOL {
	ret := C.glowDisableGenlockI3D(gpDisableGenlockI3D, (C.HDC)(hDC))
	return (C.BOOL)(ret)
}
func EnableFrameLockI3D() C.BOOL {
	ret := C.glowEnableFrameLockI3D(gpEnableFrameLockI3D)
	return (C.BOOL)(ret)
}
func EnableGenlockI3D(hDC C.HDC) C.BOOL {
	ret := C.glowEnableGenlockI3D(gpEnableGenlockI3D, (C.HDC)(hDC))
	return (C.BOOL)(ret)
}
func EndFrameTrackingI3D() C.BOOL {
	ret := C.glowEndFrameTrackingI3D(gpEndFrameTrackingI3D)
	return (C.BOOL)(ret)
}
func EnumGpuDevicesNV(hGpu C.HGPUNV, iDeviceIndex C.UINT, lpGpuDevice C.PGPU_DEVICE) C.BOOL {
	ret := C.glowEnumGpuDevicesNV(gpEnumGpuDevicesNV, (C.HGPUNV)(hGpu), (C.UINT)(iDeviceIndex), (C.PGPU_DEVICE)(lpGpuDevice))
	return (C.BOOL)(ret)
}
func EnumGpusFromAffinityDCNV(hAffinityDC C.HDC, iGpuIndex C.UINT, hGpu *C.HGPUNV) C.BOOL {
	ret := C.glowEnumGpusFromAffinityDCNV(gpEnumGpusFromAffinityDCNV, (C.HDC)(hAffinityDC), (C.UINT)(iGpuIndex), (*C.HGPUNV)(unsafe.Pointer(hGpu)))
	return (C.BOOL)(ret)
}
func EnumGpusNV(iGpuIndex C.UINT, phGpu *C.HGPUNV) C.BOOL {
	ret := C.glowEnumGpusNV(gpEnumGpusNV, (C.UINT)(iGpuIndex), (*C.HGPUNV)(unsafe.Pointer(phGpu)))
	return (C.BOOL)(ret)
}
func EnumerateVideoCaptureDevicesNV(hDc C.HDC, phDeviceList *C.HVIDEOINPUTDEVICENV) C.UINT {
	ret := C.glowEnumerateVideoCaptureDevicesNV(gpEnumerateVideoCaptureDevicesNV, (C.HDC)(hDc), (*C.HVIDEOINPUTDEVICENV)(unsafe.Pointer(phDeviceList)))
	return (C.UINT)(ret)
}
func EnumerateVideoDevicesNV(hDC C.HDC, phDeviceList *C.HVIDEOOUTPUTDEVICENV) C.int {
	ret := C.glowEnumerateVideoDevicesNV(gpEnumerateVideoDevicesNV, (C.HDC)(hDC), (*C.HVIDEOOUTPUTDEVICENV)(unsafe.Pointer(phDeviceList)))
	return (C.int)(ret)
}
func FreeMemoryNV(pointer unsafe.Pointer) {
	C.glowFreeMemoryNV(gpFreeMemoryNV, pointer)
}
func GenlockSampleRateI3D(hDC C.HDC, uRate C.UINT) C.BOOL {
	ret := C.glowGenlockSampleRateI3D(gpGenlockSampleRateI3D, (C.HDC)(hDC), (C.UINT)(uRate))
	return (C.BOOL)(ret)
}
func GenlockSourceDelayI3D(hDC C.HDC, uDelay C.UINT) C.BOOL {
	ret := C.glowGenlockSourceDelayI3D(gpGenlockSourceDelayI3D, (C.HDC)(hDC), (C.UINT)(uDelay))
	return (C.BOOL)(ret)
}
func GenlockSourceEdgeI3D(hDC C.HDC, uEdge C.UINT) C.BOOL {
	ret := C.glowGenlockSourceEdgeI3D(gpGenlockSourceEdgeI3D, (C.HDC)(hDC), (C.UINT)(uEdge))
	return (C.BOOL)(ret)
}
func GenlockSourceI3D(hDC C.HDC, uSource C.UINT) C.BOOL {
	ret := C.glowGenlockSourceI3D(gpGenlockSourceI3D, (C.HDC)(hDC), (C.UINT)(uSource))
	return (C.BOOL)(ret)
}
func GetContextGPUIDAMD(hglrc C.HGLRC) C.UINT {
	ret := C.glowGetContextGPUIDAMD(gpGetContextGPUIDAMD, (C.HGLRC)(hglrc))
	return (C.UINT)(ret)
}
func GetCurrentAssociatedContextAMD() C.HGLRC {
	ret := C.glowGetCurrentAssociatedContextAMD(gpGetCurrentAssociatedContextAMD)
	return (C.HGLRC)(ret)
}
func GetCurrentContext() C.HGLRC {
	ret := C.glowGetCurrentContext(gpGetCurrentContext)
	return (C.HGLRC)(ret)
}
func GetCurrentDC() C.HDC {
	ret := C.glowGetCurrentDC(gpGetCurrentDC)
	return (C.HDC)(ret)
}
func GetCurrentReadDCARB() C.HDC {
	ret := C.glowGetCurrentReadDCARB(gpGetCurrentReadDCARB)
	return (C.HDC)(ret)
}
func GetCurrentReadDCEXT() C.HDC {
	ret := C.glowGetCurrentReadDCEXT(gpGetCurrentReadDCEXT)
	return (C.HDC)(ret)
}
func GetDigitalVideoParametersI3D(hDC C.HDC, iAttribute C.int, piValue *C.int) C.BOOL {
	ret := C.glowGetDigitalVideoParametersI3D(gpGetDigitalVideoParametersI3D, (C.HDC)(hDC), (C.int)(iAttribute), (*C.int)(unsafe.Pointer(piValue)))
	return (C.BOOL)(ret)
}
func GetExtensionsStringARB(hdc C.HDC) *C.char {
	ret := C.glowGetExtensionsStringARB(gpGetExtensionsStringARB, (C.HDC)(hdc))
	return (*C.char)(ret)
}
func GetExtensionsStringEXT() *C.char {
	ret := C.glowGetExtensionsStringEXT(gpGetExtensionsStringEXT)
	return (*C.char)(ret)
}
func GetFrameUsageI3D(pUsage *C.float) C.BOOL {
	ret := C.glowGetFrameUsageI3D(gpGetFrameUsageI3D, (*C.float)(unsafe.Pointer(pUsage)))
	return (C.BOOL)(ret)
}
func GetGPUIDsAMD(maxCount C.UINT, ids *C.UINT) C.UINT {
	ret := C.glowGetGPUIDsAMD(gpGetGPUIDsAMD, (C.UINT)(maxCount), (*C.UINT)(unsafe.Pointer(ids)))
	return (C.UINT)(ret)
}
func GetGPUInfoAMD(id C.UINT, property C.int, dataType uint32, size C.UINT, data unsafe.Pointer) C.INT {
	ret := C.glowGetGPUInfoAMD(gpGetGPUInfoAMD, (C.UINT)(id), (C.int)(property), (C.GLenum)(dataType), (C.UINT)(size), data)
	return (C.INT)(ret)
}
func GetGammaTableI3D(hDC C.HDC, iEntries C.int, puRed *C.USHORT, puGreen *C.USHORT, puBlue *C.USHORT) C.BOOL {
	ret := C.glowGetGammaTableI3D(gpGetGammaTableI3D, (C.HDC)(hDC), (C.int)(iEntries), (*C.USHORT)(unsafe.Pointer(puRed)), (*C.USHORT)(unsafe.Pointer(puGreen)), (*C.USHORT)(unsafe.Pointer(puBlue)))
	return (C.BOOL)(ret)
}
func GetGammaTableParametersI3D(hDC C.HDC, iAttribute C.int, piValue *C.int) C.BOOL {
	ret := C.glowGetGammaTableParametersI3D(gpGetGammaTableParametersI3D, (C.HDC)(hDC), (C.int)(iAttribute), (*C.int)(unsafe.Pointer(piValue)))
	return (C.BOOL)(ret)
}
func GetGenlockSampleRateI3D(hDC C.HDC, uRate *C.UINT) C.BOOL {
	ret := C.glowGetGenlockSampleRateI3D(gpGetGenlockSampleRateI3D, (C.HDC)(hDC), (*C.UINT)(unsafe.Pointer(uRate)))
	return (C.BOOL)(ret)
}
func GetGenlockSourceDelayI3D(hDC C.HDC, uDelay *C.UINT) C.BOOL {
	ret := C.glowGetGenlockSourceDelayI3D(gpGetGenlockSourceDelayI3D, (C.HDC)(hDC), (*C.UINT)(unsafe.Pointer(uDelay)))
	return (C.BOOL)(ret)
}
func GetGenlockSourceEdgeI3D(hDC C.HDC, uEdge *C.UINT) C.BOOL {
	ret := C.glowGetGenlockSourceEdgeI3D(gpGetGenlockSourceEdgeI3D, (C.HDC)(hDC), (*C.UINT)(unsafe.Pointer(uEdge)))
	return (C.BOOL)(ret)
}
func GetGenlockSourceI3D(hDC C.HDC, uSource *C.UINT) C.BOOL {
	ret := C.glowGetGenlockSourceI3D(gpGetGenlockSourceI3D, (C.HDC)(hDC), (*C.UINT)(unsafe.Pointer(uSource)))
	return (C.BOOL)(ret)
}
func GetLayerPaletteEntries(hdc C.HDC, iLayerPlane C.int, iStart C.int, cEntries C.int, pcr *C.COLORREF) C.int {
	ret := C.glowGetLayerPaletteEntries(gpGetLayerPaletteEntries, (C.HDC)(hdc), (C.int)(iLayerPlane), (C.int)(iStart), (C.int)(cEntries), (*C.COLORREF)(unsafe.Pointer(pcr)))
	return (C.int)(ret)
}
func GetMscRateOML(hdc C.HDC, numerator *C.INT32, denominator *C.INT32) C.BOOL {
	ret := C.glowGetMscRateOML(gpGetMscRateOML, (C.HDC)(hdc), (*C.INT32)(unsafe.Pointer(numerator)), (*C.INT32)(unsafe.Pointer(denominator)))
	return (C.BOOL)(ret)
}
func GetPbufferDCARB(hPbuffer C.HPBUFFERARB) C.HDC {
	ret := C.glowGetPbufferDCARB(gpGetPbufferDCARB, (C.HPBUFFERARB)(hPbuffer))
	return (C.HDC)(ret)
}
func GetPbufferDCEXT(hPbuffer C.HPBUFFEREXT) C.HDC {
	ret := C.glowGetPbufferDCEXT(gpGetPbufferDCEXT, (C.HPBUFFEREXT)(hPbuffer))
	return (C.HDC)(ret)
}
func GetPixelFormatAttribfvARB(hdc C.HDC, iPixelFormat C.int, iLayerPlane C.int, nAttributes C.UINT, piAttributes *C.int, pfValues *C.FLOAT) C.BOOL {
	ret := C.glowGetPixelFormatAttribfvARB(gpGetPixelFormatAttribfvARB, (C.HDC)(hdc), (C.int)(iPixelFormat), (C.int)(iLayerPlane), (C.UINT)(nAttributes), (*C.int)(unsafe.Pointer(piAttributes)), (*C.FLOAT)(unsafe.Pointer(pfValues)))
	return (C.BOOL)(ret)
}
func GetPixelFormatAttribfvEXT(hdc C.HDC, iPixelFormat C.int, iLayerPlane C.int, nAttributes C.UINT, piAttributes *C.int, pfValues *C.FLOAT) C.BOOL {
	ret := C.glowGetPixelFormatAttribfvEXT(gpGetPixelFormatAttribfvEXT, (C.HDC)(hdc), (C.int)(iPixelFormat), (C.int)(iLayerPlane), (C.UINT)(nAttributes), (*C.int)(unsafe.Pointer(piAttributes)), (*C.FLOAT)(unsafe.Pointer(pfValues)))
	return (C.BOOL)(ret)
}
func GetPixelFormatAttribivARB(hdc C.HDC, iPixelFormat C.int, iLayerPlane C.int, nAttributes C.UINT, piAttributes *C.int, piValues *C.int) C.BOOL {
	ret := C.glowGetPixelFormatAttribivARB(gpGetPixelFormatAttribivARB, (C.HDC)(hdc), (C.int)(iPixelFormat), (C.int)(iLayerPlane), (C.UINT)(nAttributes), (*C.int)(unsafe.Pointer(piAttributes)), (*C.int)(unsafe.Pointer(piValues)))
	return (C.BOOL)(ret)
}
func GetPixelFormatAttribivEXT(hdc C.HDC, iPixelFormat C.int, iLayerPlane C.int, nAttributes C.UINT, piAttributes *C.int, piValues *C.int) C.BOOL {
	ret := C.glowGetPixelFormatAttribivEXT(gpGetPixelFormatAttribivEXT, (C.HDC)(hdc), (C.int)(iPixelFormat), (C.int)(iLayerPlane), (C.UINT)(nAttributes), (*C.int)(unsafe.Pointer(piAttributes)), (*C.int)(unsafe.Pointer(piValues)))
	return (C.BOOL)(ret)
}
func GetProcAddress(lpszProc C.LPCSTR) C.PROC {
	ret := C.glowGetProcAddress(gpGetProcAddress, (C.LPCSTR)(lpszProc))
	return (C.PROC)(ret)
}
func GetSwapIntervalEXT() C.int {
	ret := C.glowGetSwapIntervalEXT(gpGetSwapIntervalEXT)
	return (C.int)(ret)
}
func GetSyncValuesOML(hdc C.HDC, ust *C.INT64, msc *C.INT64, sbc *C.INT64) C.BOOL {
	ret := C.glowGetSyncValuesOML(gpGetSyncValuesOML, (C.HDC)(hdc), (*C.INT64)(unsafe.Pointer(ust)), (*C.INT64)(unsafe.Pointer(msc)), (*C.INT64)(unsafe.Pointer(sbc)))
	return (C.BOOL)(ret)
}
func GetVideoDeviceNV(hDC C.HDC, numDevices C.int, hVideoDevice *C.HPVIDEODEV) C.BOOL {
	ret := C.glowGetVideoDeviceNV(gpGetVideoDeviceNV, (C.HDC)(hDC), (C.int)(numDevices), (*C.HPVIDEODEV)(unsafe.Pointer(hVideoDevice)))
	return (C.BOOL)(ret)
}
func GetVideoInfoNV(hpVideoDevice C.HPVIDEODEV, pulCounterOutputPbuffer *C.unsignedlong, pulCounterOutputVideo *C.unsignedlong) C.BOOL {
	ret := C.glowGetVideoInfoNV(gpGetVideoInfoNV, (C.HPVIDEODEV)(hpVideoDevice), (*C.unsignedlong)(unsafe.Pointer(pulCounterOutputPbuffer)), (*C.unsignedlong)(unsafe.Pointer(pulCounterOutputVideo)))
	return (C.BOOL)(ret)
}
func IsEnabledFrameLockI3D(pFlag *C.BOOL) C.BOOL {
	ret := C.glowIsEnabledFrameLockI3D(gpIsEnabledFrameLockI3D, (*C.BOOL)(unsafe.Pointer(pFlag)))
	return (C.BOOL)(ret)
}
func IsEnabledGenlockI3D(hDC C.HDC, pFlag *C.BOOL) C.BOOL {
	ret := C.glowIsEnabledGenlockI3D(gpIsEnabledGenlockI3D, (C.HDC)(hDC), (*C.BOOL)(unsafe.Pointer(pFlag)))
	return (C.BOOL)(ret)
}
func JoinSwapGroupNV(hDC C.HDC, group uint32) C.BOOL {
	ret := C.glowJoinSwapGroupNV(gpJoinSwapGroupNV, (C.HDC)(hDC), (C.GLuint)(group))
	return (C.BOOL)(ret)
}
func LoadDisplayColorTableEXT(table *uint16, length uint32) bool {
	ret := C.glowLoadDisplayColorTableEXT(gpLoadDisplayColorTableEXT, (*C.GLushort)(unsafe.Pointer(table)), (C.GLuint)(length))
	return ret == TRUE
}
func LockVideoCaptureDeviceNV(hDc C.HDC, hDevice C.HVIDEOINPUTDEVICENV) C.BOOL {
	ret := C.glowLockVideoCaptureDeviceNV(gpLockVideoCaptureDeviceNV, (C.HDC)(hDc), (C.HVIDEOINPUTDEVICENV)(hDevice))
	return (C.BOOL)(ret)
}
func MakeAssociatedContextCurrentAMD(hglrc C.HGLRC) C.BOOL {
	ret := C.glowMakeAssociatedContextCurrentAMD(gpMakeAssociatedContextCurrentAMD, (C.HGLRC)(hglrc))
	return (C.BOOL)(ret)
}
func MakeContextCurrentARB(hDrawDC C.HDC, hReadDC C.HDC, hglrc C.HGLRC) C.BOOL {
	ret := C.glowMakeContextCurrentARB(gpMakeContextCurrentARB, (C.HDC)(hDrawDC), (C.HDC)(hReadDC), (C.HGLRC)(hglrc))
	return (C.BOOL)(ret)
}
func MakeContextCurrentEXT(hDrawDC C.HDC, hReadDC C.HDC, hglrc C.HGLRC) C.BOOL {
	ret := C.glowMakeContextCurrentEXT(gpMakeContextCurrentEXT, (C.HDC)(hDrawDC), (C.HDC)(hReadDC), (C.HGLRC)(hglrc))
	return (C.BOOL)(ret)
}
func MakeCurrent(hDc C.HDC, newContext C.HGLRC) C.BOOL {
	ret := C.glowMakeCurrent(gpMakeCurrent, (C.HDC)(hDc), (C.HGLRC)(newContext))
	return (C.BOOL)(ret)
}
func QueryCurrentContextNV(iAttribute C.int, piValue *C.int) C.BOOL {
	ret := C.glowQueryCurrentContextNV(gpQueryCurrentContextNV, (C.int)(iAttribute), (*C.int)(unsafe.Pointer(piValue)))
	return (C.BOOL)(ret)
}
func QueryFrameCountNV(hDC C.HDC, count *uint32) C.BOOL {
	ret := C.glowQueryFrameCountNV(gpQueryFrameCountNV, (C.HDC)(hDC), (*C.GLuint)(unsafe.Pointer(count)))
	return (C.BOOL)(ret)
}
func QueryFrameLockMasterI3D(pFlag *C.BOOL) C.BOOL {
	ret := C.glowQueryFrameLockMasterI3D(gpQueryFrameLockMasterI3D, (*C.BOOL)(unsafe.Pointer(pFlag)))
	return (C.BOOL)(ret)
}
func QueryFrameTrackingI3D(pFrameCount *C.DWORD, pMissedFrames *C.DWORD, pLastMissedUsage *C.float) C.BOOL {
	ret := C.glowQueryFrameTrackingI3D(gpQueryFrameTrackingI3D, (*C.DWORD)(unsafe.Pointer(pFrameCount)), (*C.DWORD)(unsafe.Pointer(pMissedFrames)), (*C.float)(unsafe.Pointer(pLastMissedUsage)))
	return (C.BOOL)(ret)
}
func QueryGenlockMaxSourceDelayI3D(hDC C.HDC, uMaxLineDelay *C.UINT, uMaxPixelDelay *C.UINT) C.BOOL {
	ret := C.glowQueryGenlockMaxSourceDelayI3D(gpQueryGenlockMaxSourceDelayI3D, (C.HDC)(hDC), (*C.UINT)(unsafe.Pointer(uMaxLineDelay)), (*C.UINT)(unsafe.Pointer(uMaxPixelDelay)))
	return (C.BOOL)(ret)
}
func QueryMaxSwapGroupsNV(hDC C.HDC, maxGroups *uint32, maxBarriers *uint32) C.BOOL {
	ret := C.glowQueryMaxSwapGroupsNV(gpQueryMaxSwapGroupsNV, (C.HDC)(hDC), (*C.GLuint)(unsafe.Pointer(maxGroups)), (*C.GLuint)(unsafe.Pointer(maxBarriers)))
	return (C.BOOL)(ret)
}
func QueryPbufferARB(hPbuffer C.HPBUFFERARB, iAttribute C.int, piValue *C.int) C.BOOL {
	ret := C.glowQueryPbufferARB(gpQueryPbufferARB, (C.HPBUFFERARB)(hPbuffer), (C.int)(iAttribute), (*C.int)(unsafe.Pointer(piValue)))
	return (C.BOOL)(ret)
}
func QueryPbufferEXT(hPbuffer C.HPBUFFEREXT, iAttribute C.int, piValue *C.int) C.BOOL {
	ret := C.glowQueryPbufferEXT(gpQueryPbufferEXT, (C.HPBUFFEREXT)(hPbuffer), (C.int)(iAttribute), (*C.int)(unsafe.Pointer(piValue)))
	return (C.BOOL)(ret)
}
func QuerySwapGroupNV(hDC C.HDC, group *uint32, barrier *uint32) C.BOOL {
	ret := C.glowQuerySwapGroupNV(gpQuerySwapGroupNV, (C.HDC)(hDC), (*C.GLuint)(unsafe.Pointer(group)), (*C.GLuint)(unsafe.Pointer(barrier)))
	return (C.BOOL)(ret)
}
func QueryVideoCaptureDeviceNV(hDc C.HDC, hDevice C.HVIDEOINPUTDEVICENV, iAttribute C.int, piValue *C.int) C.BOOL {
	ret := C.glowQueryVideoCaptureDeviceNV(gpQueryVideoCaptureDeviceNV, (C.HDC)(hDc), (C.HVIDEOINPUTDEVICENV)(hDevice), (C.int)(iAttribute), (*C.int)(unsafe.Pointer(piValue)))
	return (C.BOOL)(ret)
}
func RealizeLayerPalette(hdc C.HDC, iLayerPlane C.int, bRealize C.BOOL) C.BOOL {
	ret := C.glowRealizeLayerPalette(gpRealizeLayerPalette, (C.HDC)(hdc), (C.int)(iLayerPlane), (C.BOOL)(bRealize))
	return (C.BOOL)(ret)
}
func ReleaseImageBufferEventsI3D(hDC C.HDC, pAddress *C.LPVOID, count C.UINT) C.BOOL {
	ret := C.glowReleaseImageBufferEventsI3D(gpReleaseImageBufferEventsI3D, (C.HDC)(hDC), (*C.LPVOID)(unsafe.Pointer(pAddress)), (C.UINT)(count))
	return (C.BOOL)(ret)
}
func ReleasePbufferDCARB(hPbuffer C.HPBUFFERARB, hDC C.HDC) C.int {
	ret := C.glowReleasePbufferDCARB(gpReleasePbufferDCARB, (C.HPBUFFERARB)(hPbuffer), (C.HDC)(hDC))
	return (C.int)(ret)
}
func ReleasePbufferDCEXT(hPbuffer C.HPBUFFEREXT, hDC C.HDC) C.int {
	ret := C.glowReleasePbufferDCEXT(gpReleasePbufferDCEXT, (C.HPBUFFEREXT)(hPbuffer), (C.HDC)(hDC))
	return (C.int)(ret)
}
func ReleaseTexImageARB(hPbuffer C.HPBUFFERARB, iBuffer C.int) C.BOOL {
	ret := C.glowReleaseTexImageARB(gpReleaseTexImageARB, (C.HPBUFFERARB)(hPbuffer), (C.int)(iBuffer))
	return (C.BOOL)(ret)
}
func ReleaseVideoCaptureDeviceNV(hDc C.HDC, hDevice C.HVIDEOINPUTDEVICENV) C.BOOL {
	ret := C.glowReleaseVideoCaptureDeviceNV(gpReleaseVideoCaptureDeviceNV, (C.HDC)(hDc), (C.HVIDEOINPUTDEVICENV)(hDevice))
	return (C.BOOL)(ret)
}
func ReleaseVideoDeviceNV(hVideoDevice C.HPVIDEODEV) C.BOOL {
	ret := C.glowReleaseVideoDeviceNV(gpReleaseVideoDeviceNV, (C.HPVIDEODEV)(hVideoDevice))
	return (C.BOOL)(ret)
}
func ReleaseVideoImageNV(hPbuffer C.HPBUFFERARB, iVideoBuffer C.int) C.BOOL {
	ret := C.glowReleaseVideoImageNV(gpReleaseVideoImageNV, (C.HPBUFFERARB)(hPbuffer), (C.int)(iVideoBuffer))
	return (C.BOOL)(ret)
}
func ResetFrameCountNV(hDC C.HDC) C.BOOL {
	ret := C.glowResetFrameCountNV(gpResetFrameCountNV, (C.HDC)(hDC))
	return (C.BOOL)(ret)
}
func RestoreBufferRegionARB(hRegion C.HANDLE, x C.int, y C.int, width C.int, height C.int, xSrc C.int, ySrc C.int) C.BOOL {
	ret := C.glowRestoreBufferRegionARB(gpRestoreBufferRegionARB, (C.HANDLE)(hRegion), (C.int)(x), (C.int)(y), (C.int)(width), (C.int)(height), (C.int)(xSrc), (C.int)(ySrc))
	return (C.BOOL)(ret)
}
func SaveBufferRegionARB(hRegion C.HANDLE, x C.int, y C.int, width C.int, height C.int) C.BOOL {
	ret := C.glowSaveBufferRegionARB(gpSaveBufferRegionARB, (C.HANDLE)(hRegion), (C.int)(x), (C.int)(y), (C.int)(width), (C.int)(height))
	return (C.BOOL)(ret)
}
func SendPbufferToVideoNV(hPbuffer C.HPBUFFERARB, iBufferType C.int, pulCounterPbuffer *C.unsignedlong, bBlock C.BOOL) C.BOOL {
	ret := C.glowSendPbufferToVideoNV(gpSendPbufferToVideoNV, (C.HPBUFFERARB)(hPbuffer), (C.int)(iBufferType), (*C.unsignedlong)(unsafe.Pointer(pulCounterPbuffer)), (C.BOOL)(bBlock))
	return (C.BOOL)(ret)
}
func SetDigitalVideoParametersI3D(hDC C.HDC, iAttribute C.int, piValue *C.int) C.BOOL {
	ret := C.glowSetDigitalVideoParametersI3D(gpSetDigitalVideoParametersI3D, (C.HDC)(hDC), (C.int)(iAttribute), (*C.int)(unsafe.Pointer(piValue)))
	return (C.BOOL)(ret)
}
func SetGammaTableI3D(hDC C.HDC, iEntries C.int, puRed *C.USHORT, puGreen *C.USHORT, puBlue *C.USHORT) C.BOOL {
	ret := C.glowSetGammaTableI3D(gpSetGammaTableI3D, (C.HDC)(hDC), (C.int)(iEntries), (*C.USHORT)(unsafe.Pointer(puRed)), (*C.USHORT)(unsafe.Pointer(puGreen)), (*C.USHORT)(unsafe.Pointer(puBlue)))
	return (C.BOOL)(ret)
}
func SetGammaTableParametersI3D(hDC C.HDC, iAttribute C.int, piValue *C.int) C.BOOL {
	ret := C.glowSetGammaTableParametersI3D(gpSetGammaTableParametersI3D, (C.HDC)(hDC), (C.int)(iAttribute), (*C.int)(unsafe.Pointer(piValue)))
	return (C.BOOL)(ret)
}
func SetLayerPaletteEntries(hdc C.HDC, iLayerPlane C.int, iStart C.int, cEntries C.int, pcr *C.COLORREF) C.int {
	ret := C.glowSetLayerPaletteEntries(gpSetLayerPaletteEntries, (C.HDC)(hdc), (C.int)(iLayerPlane), (C.int)(iStart), (C.int)(cEntries), (*C.COLORREF)(unsafe.Pointer(pcr)))
	return (C.int)(ret)
}
func SetPbufferAttribARB(hPbuffer C.HPBUFFERARB, piAttribList *C.int) C.BOOL {
	ret := C.glowSetPbufferAttribARB(gpSetPbufferAttribARB, (C.HPBUFFERARB)(hPbuffer), (*C.int)(unsafe.Pointer(piAttribList)))
	return (C.BOOL)(ret)
}
func SetStereoEmitterState3DL(hDC C.HDC, uState C.UINT) C.BOOL {
	ret := C.glowSetStereoEmitterState3DL(gpSetStereoEmitterState3DL, (C.HDC)(hDC), (C.UINT)(uState))
	return (C.BOOL)(ret)
}
func ShareLists(hrcSrvShare C.HGLRC, hrcSrvSource C.HGLRC) C.BOOL {
	ret := C.glowShareLists(gpShareLists, (C.HGLRC)(hrcSrvShare), (C.HGLRC)(hrcSrvSource))
	return (C.BOOL)(ret)
}
func SwapBuffersMscOML(hdc C.HDC, target_msc C.INT64, divisor C.INT64, remainder C.INT64) C.INT64 {
	ret := C.glowSwapBuffersMscOML(gpSwapBuffersMscOML, (C.HDC)(hdc), (C.INT64)(target_msc), (C.INT64)(divisor), (C.INT64)(remainder))
	return (C.INT64)(ret)
}
func SwapIntervalEXT(interval C.int) C.BOOL {
	ret := C.glowSwapIntervalEXT(gpSwapIntervalEXT, (C.int)(interval))
	return (C.BOOL)(ret)
}
func SwapLayerBuffers(hdc C.HDC, fuFlags C.UINT) C.BOOL {
	ret := C.glowSwapLayerBuffers(gpSwapLayerBuffers, (C.HDC)(hdc), (C.UINT)(fuFlags))
	return (C.BOOL)(ret)
}
func SwapLayerBuffersMscOML(hdc C.HDC, fuPlanes C.int, target_msc C.INT64, divisor C.INT64, remainder C.INT64) C.INT64 {
	ret := C.glowSwapLayerBuffersMscOML(gpSwapLayerBuffersMscOML, (C.HDC)(hdc), (C.int)(fuPlanes), (C.INT64)(target_msc), (C.INT64)(divisor), (C.INT64)(remainder))
	return (C.INT64)(ret)
}
func UseFontBitmaps(hDC C.HDC, first C.DWORD, count C.DWORD, listBase C.DWORD) C.BOOL {
	ret := C.glowUseFontBitmaps(gpUseFontBitmaps, (C.HDC)(hDC), (C.DWORD)(first), (C.DWORD)(count), (C.DWORD)(listBase))
	return (C.BOOL)(ret)
}
func UseFontBitmapsA(hDC C.HDC, first C.DWORD, count C.DWORD, listBase C.DWORD) C.BOOL {
	ret := C.glowUseFontBitmapsA(gpUseFontBitmapsA, (C.HDC)(hDC), (C.DWORD)(first), (C.DWORD)(count), (C.DWORD)(listBase))
	return (C.BOOL)(ret)
}
func UseFontBitmapsW(hDC C.HDC, first C.DWORD, count C.DWORD, listBase C.DWORD) C.BOOL {
	ret := C.glowUseFontBitmapsW(gpUseFontBitmapsW, (C.HDC)(hDC), (C.DWORD)(first), (C.DWORD)(count), (C.DWORD)(listBase))
	return (C.BOOL)(ret)
}
func UseFontOutlines(hDC C.HDC, first C.DWORD, count C.DWORD, listBase C.DWORD, deviation C.FLOAT, extrusion C.FLOAT, format C.int, lpgmf C.LPGLYPHMETRICSFLOAT) C.BOOL {
	ret := C.glowUseFontOutlines(gpUseFontOutlines, (C.HDC)(hDC), (C.DWORD)(first), (C.DWORD)(count), (C.DWORD)(listBase), (C.FLOAT)(deviation), (C.FLOAT)(extrusion), (C.int)(format), (C.LPGLYPHMETRICSFLOAT)(lpgmf))
	return (C.BOOL)(ret)
}
func UseFontOutlinesA(hDC C.HDC, first C.DWORD, count C.DWORD, listBase C.DWORD, deviation C.FLOAT, extrusion C.FLOAT, format C.int, lpgmf C.LPGLYPHMETRICSFLOAT) C.BOOL {
	ret := C.glowUseFontOutlinesA(gpUseFontOutlinesA, (C.HDC)(hDC), (C.DWORD)(first), (C.DWORD)(count), (C.DWORD)(listBase), (C.FLOAT)(deviation), (C.FLOAT)(extrusion), (C.int)(format), (C.LPGLYPHMETRICSFLOAT)(lpgmf))
	return (C.BOOL)(ret)
}
func UseFontOutlinesW(hDC C.HDC, first C.DWORD, count C.DWORD, listBase C.DWORD, deviation C.FLOAT, extrusion C.FLOAT, format C.int, lpgmf C.LPGLYPHMETRICSFLOAT) C.BOOL {
	ret := C.glowUseFontOutlinesW(gpUseFontOutlinesW, (C.HDC)(hDC), (C.DWORD)(first), (C.DWORD)(count), (C.DWORD)(listBase), (C.FLOAT)(deviation), (C.FLOAT)(extrusion), (C.int)(format), (C.LPGLYPHMETRICSFLOAT)(lpgmf))
	return (C.BOOL)(ret)
}
func WaitForMscOML(hdc C.HDC, target_msc C.INT64, divisor C.INT64, remainder C.INT64, ust *C.INT64, msc *C.INT64, sbc *C.INT64) C.BOOL {
	ret := C.glowWaitForMscOML(gpWaitForMscOML, (C.HDC)(hdc), (C.INT64)(target_msc), (C.INT64)(divisor), (C.INT64)(remainder), (*C.INT64)(unsafe.Pointer(ust)), (*C.INT64)(unsafe.Pointer(msc)), (*C.INT64)(unsafe.Pointer(sbc)))
	return (C.BOOL)(ret)
}
func WaitForSbcOML(hdc C.HDC, target_sbc C.INT64, ust *C.INT64, msc *C.INT64, sbc *C.INT64) C.BOOL {
	ret := C.glowWaitForSbcOML(gpWaitForSbcOML, (C.HDC)(hdc), (C.INT64)(target_sbc), (*C.INT64)(unsafe.Pointer(ust)), (*C.INT64)(unsafe.Pointer(msc)), (*C.INT64)(unsafe.Pointer(sbc)))
	return (C.BOOL)(ret)
}

// Init initializes the OpenGL bindings by loading the function pointers (for
// each OpenGL function) from the active OpenGL context.
//
// It must be called under the presence of an active OpenGL context, e.g.,
// always after calling window.MakeContextCurrent() and always before calling
// any OpenGL functions exported by this package.
//
// On Windows, Init loads pointers that are context-specific (and hence you
// must re-init if switching between OpenGL contexts, although not calling Init
// again after switching between OpenGL contexts may work if the contexts belong
// to the same graphics driver/device).
//
// On macOS and the other POSIX systems, the behavior is different, but code
// written compatible with the Windows behavior is compatible with macOS and the
// other POSIX systems. That is, always Init under an active OpenGL context, and
// always re-init after switching graphics contexts.
//
// For information about caveats of Init, you should read the "Platform Specific
// Function Retrieval" section of https://www.opengl.org/wiki/Load_OpenGL_Functions.
func Init() error {
	return InitWithProcAddrFunc(getProcAddress)
}

// InitWithProcAddrFunc intializes the package using the specified OpenGL
// function pointer loading function. For more cases Init should be used
// instead.
func InitWithProcAddrFunc(getProcAddr func(name string) unsafe.Pointer) error {
	gpChoosePixelFormat = (C.GPCHOOSEPIXELFORMAT)(getProcAddr("ChoosePixelFormat"))
	if gpChoosePixelFormat == nil {
		return errors.New("ChoosePixelFormat")
	}
	gpDescribePixelFormat = (C.GPDESCRIBEPIXELFORMAT)(getProcAddr("DescribePixelFormat"))
	if gpDescribePixelFormat == nil {
		return errors.New("DescribePixelFormat")
	}
	gpGetEnhMetaFilePixelFormat = (C.GPGETENHMETAFILEPIXELFORMAT)(getProcAddr("GetEnhMetaFilePixelFormat"))
	if gpGetEnhMetaFilePixelFormat == nil {
		return errors.New("GetEnhMetaFilePixelFormat")
	}
	gpGetPixelFormat = (C.GPGETPIXELFORMAT)(getProcAddr("GetPixelFormat"))
	if gpGetPixelFormat == nil {
		return errors.New("GetPixelFormat")
	}
	gpSetPixelFormat = (C.GPSETPIXELFORMAT)(getProcAddr("SetPixelFormat"))
	if gpSetPixelFormat == nil {
		return errors.New("SetPixelFormat")
	}
	gpSwapBuffers = (C.GPSWAPBUFFERS)(getProcAddr("SwapBuffers"))
	if gpSwapBuffers == nil {
		return errors.New("SwapBuffers")
	}
	gpAllocateMemoryNV = (C.GPALLOCATEMEMORYNV)(getProcAddr("wglAllocateMemoryNV"))
	gpAssociateImageBufferEventsI3D = (C.GPASSOCIATEIMAGEBUFFEREVENTSI3D)(getProcAddr("wglAssociateImageBufferEventsI3D"))
	gpBeginFrameTrackingI3D = (C.GPBEGINFRAMETRACKINGI3D)(getProcAddr("wglBeginFrameTrackingI3D"))
	gpBindDisplayColorTableEXT = (C.GPBINDDISPLAYCOLORTABLEEXT)(getProcAddr("wglBindDisplayColorTableEXT"))
	gpBindSwapBarrierNV = (C.GPBINDSWAPBARRIERNV)(getProcAddr("wglBindSwapBarrierNV"))
	gpBindTexImageARB = (C.GPBINDTEXIMAGEARB)(getProcAddr("wglBindTexImageARB"))
	gpBindVideoCaptureDeviceNV = (C.GPBINDVIDEOCAPTUREDEVICENV)(getProcAddr("wglBindVideoCaptureDeviceNV"))
	gpBindVideoDeviceNV = (C.GPBINDVIDEODEVICENV)(getProcAddr("wglBindVideoDeviceNV"))
	gpBindVideoImageNV = (C.GPBINDVIDEOIMAGENV)(getProcAddr("wglBindVideoImageNV"))
	gpBlitContextFramebufferAMD = (C.GPBLITCONTEXTFRAMEBUFFERAMD)(getProcAddr("wglBlitContextFramebufferAMD"))
	gpChoosePixelFormatARB = (C.GPCHOOSEPIXELFORMATARB)(getProcAddr("wglChoosePixelFormatARB"))
	gpChoosePixelFormatEXT = (C.GPCHOOSEPIXELFORMATEXT)(getProcAddr("wglChoosePixelFormatEXT"))
	gpCopyContext = (C.GPCOPYCONTEXT)(getProcAddr("wglCopyContext"))
	if gpCopyContext == nil {
		return errors.New("wglCopyContext")
	}
	gpCopyImageSubDataNV = (C.GPCOPYIMAGESUBDATANV)(getProcAddr("wglCopyImageSubDataNV"))
	gpCreateAffinityDCNV = (C.GPCREATEAFFINITYDCNV)(getProcAddr("wglCreateAffinityDCNV"))
	gpCreateAssociatedContextAMD = (C.GPCREATEASSOCIATEDCONTEXTAMD)(getProcAddr("wglCreateAssociatedContextAMD"))
	gpCreateAssociatedContextAttribsAMD = (C.GPCREATEASSOCIATEDCONTEXTATTRIBSAMD)(getProcAddr("wglCreateAssociatedContextAttribsAMD"))
	gpCreateBufferRegionARB = (C.GPCREATEBUFFERREGIONARB)(getProcAddr("wglCreateBufferRegionARB"))
	gpCreateContext = (C.GPCREATECONTEXT)(getProcAddr("wglCreateContext"))
	if gpCreateContext == nil {
		return errors.New("wglCreateContext")
	}
	gpCreateContextAttribsARB = (C.GPCREATECONTEXTATTRIBSARB)(getProcAddr("wglCreateContextAttribsARB"))
	gpCreateDisplayColorTableEXT = (C.GPCREATEDISPLAYCOLORTABLEEXT)(getProcAddr("wglCreateDisplayColorTableEXT"))
	gpCreateImageBufferI3D = (C.GPCREATEIMAGEBUFFERI3D)(getProcAddr("wglCreateImageBufferI3D"))
	gpCreateLayerContext = (C.GPCREATELAYERCONTEXT)(getProcAddr("wglCreateLayerContext"))
	if gpCreateLayerContext == nil {
		return errors.New("wglCreateLayerContext")
	}
	gpCreatePbufferARB = (C.GPCREATEPBUFFERARB)(getProcAddr("wglCreatePbufferARB"))
	gpCreatePbufferEXT = (C.GPCREATEPBUFFEREXT)(getProcAddr("wglCreatePbufferEXT"))
	gpDXCloseDeviceNV = (C.GPDXCLOSEDEVICENV)(getProcAddr("wglDXCloseDeviceNV"))
	gpDXLockObjectsNV = (C.GPDXLOCKOBJECTSNV)(getProcAddr("wglDXLockObjectsNV"))
	gpDXObjectAccessNV = (C.GPDXOBJECTACCESSNV)(getProcAddr("wglDXObjectAccessNV"))
	gpDXOpenDeviceNV = (C.GPDXOPENDEVICENV)(getProcAddr("wglDXOpenDeviceNV"))
	gpDXRegisterObjectNV = (C.GPDXREGISTEROBJECTNV)(getProcAddr("wglDXRegisterObjectNV"))
	gpDXSetResourceShareHandleNV = (C.GPDXSETRESOURCESHAREHANDLENV)(getProcAddr("wglDXSetResourceShareHandleNV"))
	gpDXUnlockObjectsNV = (C.GPDXUNLOCKOBJECTSNV)(getProcAddr("wglDXUnlockObjectsNV"))
	gpDXUnregisterObjectNV = (C.GPDXUNREGISTEROBJECTNV)(getProcAddr("wglDXUnregisterObjectNV"))
	gpDelayBeforeSwapNV = (C.GPDELAYBEFORESWAPNV)(getProcAddr("wglDelayBeforeSwapNV"))
	gpDeleteAssociatedContextAMD = (C.GPDELETEASSOCIATEDCONTEXTAMD)(getProcAddr("wglDeleteAssociatedContextAMD"))
	gpDeleteBufferRegionARB = (C.GPDELETEBUFFERREGIONARB)(getProcAddr("wglDeleteBufferRegionARB"))
	gpDeleteContext = (C.GPDELETECONTEXT)(getProcAddr("wglDeleteContext"))
	if gpDeleteContext == nil {
		return errors.New("wglDeleteContext")
	}
	gpDeleteDCNV = (C.GPDELETEDCNV)(getProcAddr("wglDeleteDCNV"))
	gpDescribeLayerPlane = (C.GPDESCRIBELAYERPLANE)(getProcAddr("wglDescribeLayerPlane"))
	if gpDescribeLayerPlane == nil {
		return errors.New("wglDescribeLayerPlane")
	}
	gpDestroyDisplayColorTableEXT = (C.GPDESTROYDISPLAYCOLORTABLEEXT)(getProcAddr("wglDestroyDisplayColorTableEXT"))
	gpDestroyImageBufferI3D = (C.GPDESTROYIMAGEBUFFERI3D)(getProcAddr("wglDestroyImageBufferI3D"))
	gpDestroyPbufferARB = (C.GPDESTROYPBUFFERARB)(getProcAddr("wglDestroyPbufferARB"))
	gpDestroyPbufferEXT = (C.GPDESTROYPBUFFEREXT)(getProcAddr("wglDestroyPbufferEXT"))
	gpDisableFrameLockI3D = (C.GPDISABLEFRAMELOCKI3D)(getProcAddr("wglDisableFrameLockI3D"))
	gpDisableGenlockI3D = (C.GPDISABLEGENLOCKI3D)(getProcAddr("wglDisableGenlockI3D"))
	gpEnableFrameLockI3D = (C.GPENABLEFRAMELOCKI3D)(getProcAddr("wglEnableFrameLockI3D"))
	gpEnableGenlockI3D = (C.GPENABLEGENLOCKI3D)(getProcAddr("wglEnableGenlockI3D"))
	gpEndFrameTrackingI3D = (C.GPENDFRAMETRACKINGI3D)(getProcAddr("wglEndFrameTrackingI3D"))
	gpEnumGpuDevicesNV = (C.GPENUMGPUDEVICESNV)(getProcAddr("wglEnumGpuDevicesNV"))
	gpEnumGpusFromAffinityDCNV = (C.GPENUMGPUSFROMAFFINITYDCNV)(getProcAddr("wglEnumGpusFromAffinityDCNV"))
	gpEnumGpusNV = (C.GPENUMGPUSNV)(getProcAddr("wglEnumGpusNV"))
	gpEnumerateVideoCaptureDevicesNV = (C.GPENUMERATEVIDEOCAPTUREDEVICESNV)(getProcAddr("wglEnumerateVideoCaptureDevicesNV"))
	gpEnumerateVideoDevicesNV = (C.GPENUMERATEVIDEODEVICESNV)(getProcAddr("wglEnumerateVideoDevicesNV"))
	gpFreeMemoryNV = (C.GPFREEMEMORYNV)(getProcAddr("wglFreeMemoryNV"))
	gpGenlockSampleRateI3D = (C.GPGENLOCKSAMPLERATEI3D)(getProcAddr("wglGenlockSampleRateI3D"))
	gpGenlockSourceDelayI3D = (C.GPGENLOCKSOURCEDELAYI3D)(getProcAddr("wglGenlockSourceDelayI3D"))
	gpGenlockSourceEdgeI3D = (C.GPGENLOCKSOURCEEDGEI3D)(getProcAddr("wglGenlockSourceEdgeI3D"))
	gpGenlockSourceI3D = (C.GPGENLOCKSOURCEI3D)(getProcAddr("wglGenlockSourceI3D"))
	gpGetContextGPUIDAMD = (C.GPGETCONTEXTGPUIDAMD)(getProcAddr("wglGetContextGPUIDAMD"))
	gpGetCurrentAssociatedContextAMD = (C.GPGETCURRENTASSOCIATEDCONTEXTAMD)(getProcAddr("wglGetCurrentAssociatedContextAMD"))
	gpGetCurrentContext = (C.GPGETCURRENTCONTEXT)(getProcAddr("wglGetCurrentContext"))
	if gpGetCurrentContext == nil {
		return errors.New("wglGetCurrentContext")
	}
	gpGetCurrentDC = (C.GPGETCURRENTDC)(getProcAddr("wglGetCurrentDC"))
	if gpGetCurrentDC == nil {
		return errors.New("wglGetCurrentDC")
	}
	gpGetCurrentReadDCARB = (C.GPGETCURRENTREADDCARB)(getProcAddr("wglGetCurrentReadDCARB"))
	gpGetCurrentReadDCEXT = (C.GPGETCURRENTREADDCEXT)(getProcAddr("wglGetCurrentReadDCEXT"))
	gpGetDigitalVideoParametersI3D = (C.GPGETDIGITALVIDEOPARAMETERSI3D)(getProcAddr("wglGetDigitalVideoParametersI3D"))
	gpGetExtensionsStringARB = (C.GPGETEXTENSIONSSTRINGARB)(getProcAddr("wglGetExtensionsStringARB"))
	gpGetExtensionsStringEXT = (C.GPGETEXTENSIONSSTRINGEXT)(getProcAddr("wglGetExtensionsStringEXT"))
	gpGetFrameUsageI3D = (C.GPGETFRAMEUSAGEI3D)(getProcAddr("wglGetFrameUsageI3D"))
	gpGetGPUIDsAMD = (C.GPGETGPUIDSAMD)(getProcAddr("wglGetGPUIDsAMD"))
	gpGetGPUInfoAMD = (C.GPGETGPUINFOAMD)(getProcAddr("wglGetGPUInfoAMD"))
	gpGetGammaTableI3D = (C.GPGETGAMMATABLEI3D)(getProcAddr("wglGetGammaTableI3D"))
	gpGetGammaTableParametersI3D = (C.GPGETGAMMATABLEPARAMETERSI3D)(getProcAddr("wglGetGammaTableParametersI3D"))
	gpGetGenlockSampleRateI3D = (C.GPGETGENLOCKSAMPLERATEI3D)(getProcAddr("wglGetGenlockSampleRateI3D"))
	gpGetGenlockSourceDelayI3D = (C.GPGETGENLOCKSOURCEDELAYI3D)(getProcAddr("wglGetGenlockSourceDelayI3D"))
	gpGetGenlockSourceEdgeI3D = (C.GPGETGENLOCKSOURCEEDGEI3D)(getProcAddr("wglGetGenlockSourceEdgeI3D"))
	gpGetGenlockSourceI3D = (C.GPGETGENLOCKSOURCEI3D)(getProcAddr("wglGetGenlockSourceI3D"))
	gpGetLayerPaletteEntries = (C.GPGETLAYERPALETTEENTRIES)(getProcAddr("wglGetLayerPaletteEntries"))
	if gpGetLayerPaletteEntries == nil {
		return errors.New("wglGetLayerPaletteEntries")
	}
	gpGetMscRateOML = (C.GPGETMSCRATEOML)(getProcAddr("wglGetMscRateOML"))
	gpGetPbufferDCARB = (C.GPGETPBUFFERDCARB)(getProcAddr("wglGetPbufferDCARB"))
	gpGetPbufferDCEXT = (C.GPGETPBUFFERDCEXT)(getProcAddr("wglGetPbufferDCEXT"))
	gpGetPixelFormatAttribfvARB = (C.GPGETPIXELFORMATATTRIBFVARB)(getProcAddr("wglGetPixelFormatAttribfvARB"))
	gpGetPixelFormatAttribfvEXT = (C.GPGETPIXELFORMATATTRIBFVEXT)(getProcAddr("wglGetPixelFormatAttribfvEXT"))
	gpGetPixelFormatAttribivARB = (C.GPGETPIXELFORMATATTRIBIVARB)(getProcAddr("wglGetPixelFormatAttribivARB"))
	gpGetPixelFormatAttribivEXT = (C.GPGETPIXELFORMATATTRIBIVEXT)(getProcAddr("wglGetPixelFormatAttribivEXT"))
	gpGetProcAddress = (C.GPGETPROCADDRESS)(getProcAddr("wglGetProcAddress"))
	if gpGetProcAddress == nil {
		return errors.New("wglGetProcAddress")
	}
	gpGetSwapIntervalEXT = (C.GPGETSWAPINTERVALEXT)(getProcAddr("wglGetSwapIntervalEXT"))
	gpGetSyncValuesOML = (C.GPGETSYNCVALUESOML)(getProcAddr("wglGetSyncValuesOML"))
	gpGetVideoDeviceNV = (C.GPGETVIDEODEVICENV)(getProcAddr("wglGetVideoDeviceNV"))
	gpGetVideoInfoNV = (C.GPGETVIDEOINFONV)(getProcAddr("wglGetVideoInfoNV"))
	gpIsEnabledFrameLockI3D = (C.GPISENABLEDFRAMELOCKI3D)(getProcAddr("wglIsEnabledFrameLockI3D"))
	gpIsEnabledGenlockI3D = (C.GPISENABLEDGENLOCKI3D)(getProcAddr("wglIsEnabledGenlockI3D"))
	gpJoinSwapGroupNV = (C.GPJOINSWAPGROUPNV)(getProcAddr("wglJoinSwapGroupNV"))
	gpLoadDisplayColorTableEXT = (C.GPLOADDISPLAYCOLORTABLEEXT)(getProcAddr("wglLoadDisplayColorTableEXT"))
	gpLockVideoCaptureDeviceNV = (C.GPLOCKVIDEOCAPTUREDEVICENV)(getProcAddr("wglLockVideoCaptureDeviceNV"))
	gpMakeAssociatedContextCurrentAMD = (C.GPMAKEASSOCIATEDCONTEXTCURRENTAMD)(getProcAddr("wglMakeAssociatedContextCurrentAMD"))
	gpMakeContextCurrentARB = (C.GPMAKECONTEXTCURRENTARB)(getProcAddr("wglMakeContextCurrentARB"))
	gpMakeContextCurrentEXT = (C.GPMAKECONTEXTCURRENTEXT)(getProcAddr("wglMakeContextCurrentEXT"))
	gpMakeCurrent = (C.GPMAKECURRENT)(getProcAddr("wglMakeCurrent"))
	if gpMakeCurrent == nil {
		return errors.New("wglMakeCurrent")
	}
	gpQueryCurrentContextNV = (C.GPQUERYCURRENTCONTEXTNV)(getProcAddr("wglQueryCurrentContextNV"))
	gpQueryFrameCountNV = (C.GPQUERYFRAMECOUNTNV)(getProcAddr("wglQueryFrameCountNV"))
	gpQueryFrameLockMasterI3D = (C.GPQUERYFRAMELOCKMASTERI3D)(getProcAddr("wglQueryFrameLockMasterI3D"))
	gpQueryFrameTrackingI3D = (C.GPQUERYFRAMETRACKINGI3D)(getProcAddr("wglQueryFrameTrackingI3D"))
	gpQueryGenlockMaxSourceDelayI3D = (C.GPQUERYGENLOCKMAXSOURCEDELAYI3D)(getProcAddr("wglQueryGenlockMaxSourceDelayI3D"))
	gpQueryMaxSwapGroupsNV = (C.GPQUERYMAXSWAPGROUPSNV)(getProcAddr("wglQueryMaxSwapGroupsNV"))
	gpQueryPbufferARB = (C.GPQUERYPBUFFERARB)(getProcAddr("wglQueryPbufferARB"))
	gpQueryPbufferEXT = (C.GPQUERYPBUFFEREXT)(getProcAddr("wglQueryPbufferEXT"))
	gpQuerySwapGroupNV = (C.GPQUERYSWAPGROUPNV)(getProcAddr("wglQuerySwapGroupNV"))
	gpQueryVideoCaptureDeviceNV = (C.GPQUERYVIDEOCAPTUREDEVICENV)(getProcAddr("wglQueryVideoCaptureDeviceNV"))
	gpRealizeLayerPalette = (C.GPREALIZELAYERPALETTE)(getProcAddr("wglRealizeLayerPalette"))
	if gpRealizeLayerPalette == nil {
		return errors.New("wglRealizeLayerPalette")
	}
	gpReleaseImageBufferEventsI3D = (C.GPRELEASEIMAGEBUFFEREVENTSI3D)(getProcAddr("wglReleaseImageBufferEventsI3D"))
	gpReleasePbufferDCARB = (C.GPRELEASEPBUFFERDCARB)(getProcAddr("wglReleasePbufferDCARB"))
	gpReleasePbufferDCEXT = (C.GPRELEASEPBUFFERDCEXT)(getProcAddr("wglReleasePbufferDCEXT"))
	gpReleaseTexImageARB = (C.GPRELEASETEXIMAGEARB)(getProcAddr("wglReleaseTexImageARB"))
	gpReleaseVideoCaptureDeviceNV = (C.GPRELEASEVIDEOCAPTUREDEVICENV)(getProcAddr("wglReleaseVideoCaptureDeviceNV"))
	gpReleaseVideoDeviceNV = (C.GPRELEASEVIDEODEVICENV)(getProcAddr("wglReleaseVideoDeviceNV"))
	gpReleaseVideoImageNV = (C.GPRELEASEVIDEOIMAGENV)(getProcAddr("wglReleaseVideoImageNV"))
	gpResetFrameCountNV = (C.GPRESETFRAMECOUNTNV)(getProcAddr("wglResetFrameCountNV"))
	gpRestoreBufferRegionARB = (C.GPRESTOREBUFFERREGIONARB)(getProcAddr("wglRestoreBufferRegionARB"))
	gpSaveBufferRegionARB = (C.GPSAVEBUFFERREGIONARB)(getProcAddr("wglSaveBufferRegionARB"))
	gpSendPbufferToVideoNV = (C.GPSENDPBUFFERTOVIDEONV)(getProcAddr("wglSendPbufferToVideoNV"))
	gpSetDigitalVideoParametersI3D = (C.GPSETDIGITALVIDEOPARAMETERSI3D)(getProcAddr("wglSetDigitalVideoParametersI3D"))
	gpSetGammaTableI3D = (C.GPSETGAMMATABLEI3D)(getProcAddr("wglSetGammaTableI3D"))
	gpSetGammaTableParametersI3D = (C.GPSETGAMMATABLEPARAMETERSI3D)(getProcAddr("wglSetGammaTableParametersI3D"))
	gpSetLayerPaletteEntries = (C.GPSETLAYERPALETTEENTRIES)(getProcAddr("wglSetLayerPaletteEntries"))
	if gpSetLayerPaletteEntries == nil {
		return errors.New("wglSetLayerPaletteEntries")
	}
	gpSetPbufferAttribARB = (C.GPSETPBUFFERATTRIBARB)(getProcAddr("wglSetPbufferAttribARB"))
	gpSetStereoEmitterState3DL = (C.GPSETSTEREOEMITTERSTATE3DL)(getProcAddr("wglSetStereoEmitterState3DL"))
	gpShareLists = (C.GPSHARELISTS)(getProcAddr("wglShareLists"))
	if gpShareLists == nil {
		return errors.New("wglShareLists")
	}
	gpSwapBuffersMscOML = (C.GPSWAPBUFFERSMSCOML)(getProcAddr("wglSwapBuffersMscOML"))
	gpSwapIntervalEXT = (C.GPSWAPINTERVALEXT)(getProcAddr("wglSwapIntervalEXT"))
	gpSwapLayerBuffers = (C.GPSWAPLAYERBUFFERS)(getProcAddr("wglSwapLayerBuffers"))
	if gpSwapLayerBuffers == nil {
		return errors.New("wglSwapLayerBuffers")
	}
	gpSwapLayerBuffersMscOML = (C.GPSWAPLAYERBUFFERSMSCOML)(getProcAddr("wglSwapLayerBuffersMscOML"))
	gpUseFontBitmaps = (C.GPUSEFONTBITMAPS)(getProcAddr("wglUseFontBitmaps"))
	if gpUseFontBitmaps == nil {
		return errors.New("wglUseFontBitmaps")
	}
	gpUseFontBitmapsA = (C.GPUSEFONTBITMAPSA)(getProcAddr("wglUseFontBitmapsA"))
	if gpUseFontBitmapsA == nil {
		return errors.New("wglUseFontBitmapsA")
	}
	gpUseFontBitmapsW = (C.GPUSEFONTBITMAPSW)(getProcAddr("wglUseFontBitmapsW"))
	if gpUseFontBitmapsW == nil {
		return errors.New("wglUseFontBitmapsW")
	}
	gpUseFontOutlines = (C.GPUSEFONTOUTLINES)(getProcAddr("wglUseFontOutlines"))
	if gpUseFontOutlines == nil {
		return errors.New("wglUseFontOutlines")
	}
	gpUseFontOutlinesA = (C.GPUSEFONTOUTLINESA)(getProcAddr("wglUseFontOutlinesA"))
	if gpUseFontOutlinesA == nil {
		return errors.New("wglUseFontOutlinesA")
	}
	gpUseFontOutlinesW = (C.GPUSEFONTOUTLINESW)(getProcAddr("wglUseFontOutlinesW"))
	if gpUseFontOutlinesW == nil {
		return errors.New("wglUseFontOutlinesW")
	}
	gpWaitForMscOML = (C.GPWAITFORMSCOML)(getProcAddr("wglWaitForMscOML"))
	gpWaitForSbcOML = (C.GPWAITFORSBCOML)(getProcAddr("wglWaitForSbcOML"))
	return nil
}
