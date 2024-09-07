Name:           MineServerTools
Version:        1.1.1
Release:        1%{?dist}
Summary:        Ferramentas essenciais para manter um servidor Minecraft sempre ativo

License:        GPL
Source0:        %{name}-%{version}.tar.gz
BuildArch:      noarch

Requires:       tar
Requires:       wget
Requires:       rsync
Requires:       tmux
Requires:       unzip
Requires:       lsof
Requires:       dialog

%description
Descrição detalhada do pacote.

%prep
%setup -q

%build

%install
mkdir -p %{buildroot}/etc/mineservertools
cp -r etc/mineservertools/* %{buildroot}/etc/mineservertools/
mkdir -p %{buildroot}/usr/bin
cp -r usr/bin/* %{buildroot}/usr/bin/
mkdir -p %{buildroot}/lib/systemd/system
cp -r lib/systemd/system/* %{buildroot}/lib/systemd/system/
mkdir -p %{buildroot}/var/log
cp -r var/log/* %{buildroot}/var/log/
mkdir -p %{buildroot}/var/mine-backups
cp -r var/mine-backups/* %{buildroot}/var/mine-backups/

%post
# Executar script pós-instalação
/bin/bash /etc/mine-server-tools/post-install.sh

%preun
# Comandos de pré-desinstalação (se necessário)

%postun
# Comandos de pós-desinstalação (se necessário)

%files
/etc/mineservertools/*
/usr/bin/*
/lib/systemd/system/*
/var/log/*
/var/mine-backups/*
