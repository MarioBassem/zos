YGGDRASIL_VERSION="v0.4.0"
YGGDRASIL_CHECKSUM="5f5d58c904392bd892816f5ba793c85f"
YGGDRASIL_LINK="https://github.com/yggdrasil-network/yggdrasil-go/releases/download/v${YGGDRASIL_VERSION}/yggdrasil-${YGGDRASIL_VERSION}-amd64.deb"

download_yggdrasil() {
    download_file ${YGGDRASIL_LINK} ${YGGDRASIL_CHECKSUM} yggdrasil-v${YGGDRASIL_VERSION}.deb
}

extract_yggdrasil() {
    apt-get install ./yggdrasil-v${YGGDRASIL_VERSION}.deb
}

prepare_yggdrasil() {
    echo "[+] prepare yggdrasil"
    github_name "yggdrasil-${YGGDRASIL_VERSION}"
}

install_yggdrasil() {
    echo "[+] install yggdrasil"

    mkdir -p "${ROOTDIR}/usr/bin"
    mkdir -p "${ROOTDIR}/etc/yggdrasil"
    mkdir -p "${ROOTDIR}/etc/zinit"

    cp -av $(which yggdrasil) "${ROOTDIR}/usr/bin/"
    cp -av $(which yggdrasilctl) "${ROOTDIR}/usr/bin/"
}

build_yggdrasil() {
    pushd "${DISTDIR}"
    download_yggdrasil
    extract_yggdrasil
    popd

    prepare_yggdrasil
    install_yggdrasil
}
