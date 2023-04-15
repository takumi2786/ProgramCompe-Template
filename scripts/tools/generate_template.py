# golangのテンプレートをコピーする
# vscode > launchから実行する想定
import argparse
import os
import glob
import shutil


BASE_DIR = os.path.abspath(os.path.join(os.path.dirname(__file__), ".."))
TEMPLATE_DIR = os.path.join(BASE_DIR, "resources/templates")
PROGRAM_TYPES = ['python', 'golang', 'java']

def main(program_type: str, out_dir: str):
    # ディレクトリ以下のファイルをコピーする
    if program_type not in PROGRAM_TYPES:
        raise Exception(f'不正なプログラムタイプです: {program_type}')
    input_paths = glob.glob( os.path.join(TEMPLATE_DIR, program_type, "*") )
    for path in input_paths:
        file_name = os.path.basename(path)
        try:
            shutil.copyfile(path, os.path.join(out_dir, file_name))
        except IsADirectoryError:
            shutil.copytree(path, os.path.join(out_dir, file_name))
            continue


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('-t', '--type', type=str, default=None, help='python or golang')
    parser.add_argument('-d', '--dirname', type=str, default=None, help='テンプレートを配置するディレクトリ')
    args = parser.parse_args()
    program_type = args.type
    dir_name = args.dirname
    main(program_type, dir_name)
