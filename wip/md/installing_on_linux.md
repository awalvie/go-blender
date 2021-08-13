---
title: installing things
subtitle: linux be weird
created: 2020-0-04
---

I was trying to install the Godot Engine for developing a game. I went to their [website](https://godotengine.org/download/linux), downloaded and extracted it and it gave me a single file executable. I ran it and it worked, so why the post?

Well, if you've been on linux for long you'll know that usually you install things from the terminal (with a package manager often) and then run them, you don't have to deal with manually placing binaries or executables in the right place. So I thought it would be a good opportunity to share how I learnt to do it. In the case of Godot, these were the steps I followed:

Download Godot, unzip it using:

```sh
unzip Godot*.zip
```

Move the executable to `/opt`:

```sh
sudo mkdir /opt/Godot/
sudo mv Godot_*.64 /opt/Godot/Godot
```

Make it executable:

```sh
sudo chmod +x /opt/Godot/Godot
```

Now we create a launcher for the Godot executable:

```sh
cd ~/.local/share/applications/
vim godot.desktop
```

Populate the `.desktop` file, writing this in nano, saving with `CTRL-O` and exiting with `CTRL-X`:

```sh
[Desktop Entry]
Name=Godot Engine

GenericName=Libre game engine
Comment=Multi-platform 2D and 3D game engine with a feature rich editor
Exec=/opt/Godot/Godot -pm
Icon=godot
Terminal=false
Type=Application
Categories=Development;IDE;
```

and voila, you have _installed_ Godot, this should be applicable for most other installs as well.