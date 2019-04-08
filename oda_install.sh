# ダウンロード
curl -O http://dspcdn.sonycreativesoftware.com/partners/OpticalDiscArchiveSoftware/ODASoftware_4_4_0_linux.zip

# unzip がなければ、
# yum install -y unzip

unzip ODASoftware_4_4_0_linux.zip

# 依存関係インストール（which はインストール時の警告が出ないようにするため。必須ではない。）
yum install -y which lsof fuse-libs apr apr-util

./ODASoftware_V440.bin

oda-util --version
