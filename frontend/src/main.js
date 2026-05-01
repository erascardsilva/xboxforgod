import {
    VerificarDependencias,
    ListarDispositivosCDROM,
    SelecionarDiretorio,
    SelecionarArquivoISO,
    CopiarISO,
    ConverterISOparaGOD,
    GetPlatform
} from '../wailsjs/go/main/App';

import { EventsOn, BrowserOpenURL } from '../wailsjs/runtime/runtime';

const translations = {
    pt: {
        nav_convert_label: "Converter",
        nav_convert_btn: "Converter ISO para GOD",
        nav_copy_btn: "Cópia DVD para ISO",
        nav_about_btn: "Sobre",
        nav_support_btn: "Apoie o Projeto",
        status_checking: "Verificando...",
        status_deps_ok: "Dependências OK",
        status_missing: "Faltando: ",
        status_err_verify: "Erro ao verificar",
        page_convert_title: "ISO → GOD",
        page_convert_desc: "Transforme arquivos ISO em formato GOD para uso em Xbox 360 (RGH/JTAG)",
        label_iso_input: "Arquivo ISO de entrada",
        placeholder_select_iso: "Selecione o arquivo ISO...",
        btn_browse: "Procurar",
        label_god_output: "Diretório de saída GOD",
        placeholder_select_dir: "Selecione o diretório...",
        btn_convert_action: "Converter para GOD",
        progress_waiting: "Aguardando...",
        page_copy_title: "Copiar ISO de DVD",
        page_copy_desc: "Crie uma imagem ISO a partir do seu DVD de jogo Xbox 360",
        label_dvd_drive: "Unidade de DVD",
        select_drive_default: "Selecione uma unidade...",
        select_drive_searching: "Buscando unidades...",
        select_drive_empty: "Nenhuma unidade óptica encontrada",
        select_drive_err: "Erro ao listar unidades",
        label_output_dir: "Diretório de saída",
        label_iso_name: "Nome do arquivo ISO",
        placeholder_iso_name: "meu_jogo.iso",
        btn_copy_action: "Iniciar Cópia",
        about_title: "Sobre o XboxForGOD",
        about_desc: "Como funciona e para que serve",
        about_q1: "O que é o formato GOD?",
        about_a1: "O formato <strong>GOD (Games on Demand)</strong> é um formato de arquivo específico para jogos do Xbox 360, projetado para ser executado diretamente do disco rígido do console. É amplamente utilizado em consoles Xbox 360 modificados com <strong>RGH/JTAG</strong>.",
        about_q2: "Como funciona?",
        about_step1_title: "Inserir DVD",
        about_step1_desc: "Insira o DVD do jogo Xbox 360 no leitor óptico do computador.",
        about_step2_title: "Copiar ISO",
        about_step2_desc: "O programa utiliza o comando <code>dd</code> (Linux) ou leitura direta (Windows) para criar uma imagem ISO idêntica do disco.",
        about_step3_title: "Converter para GOD",
        about_step3_desc: "A ISO é convertida para o formato GOD usando o utilitário <code>iso2god</code>, gerando a estrutura de pastas compatível com o Xbox 360.",
        about_step4_title: "Copiar para HD",
        about_step4_desc: "A pasta GOD gerada é copiada para o HD interno ou externo do Xbox 360 com exploit (RGH/JTAG), aparecendo diretamente no menu de jogos.",
        about_adv_title: "Vantagens do formato GOD",
        about_adv1: "<strong>Conveniência</strong> — Jogos executados diretamente do HD, sem disco físico.",
        about_adv2: "<strong>Velocidade</strong> — Carregamento mais rápido que a leitura do DVD.",
        about_adv3: "<strong>Preservação</strong> — Reduz o desgaste do leitor de DVD do console.",
        about_adv4: "<strong>Organização</strong> — Facilita o gerenciamento de grandes bibliotecas de jogos.",
        about_req_title: "Requisitos do sistema",
        about_req_detected: "Plataforma detectada",
        about_dev_by: "Desenvolvido por",
        
        log_err_dir: "Erro ao selecionar diretório: ",
        log_err_file: "Erro ao selecionar arquivo: ",
        log_warn_dvd: "Selecione uma unidade de DVD.",
        log_warn_dir: "Selecione um diretório de saída.",
        log_warn_name: "Digite um nome para o arquivo ISO.",
        log_warn_iso: "Selecione um arquivo ISO.",
        btn_copying: "Copiando...",
        btn_converting: "Convertendo...",
        log_done: "Concluído: ",
        log_error: "Erro: ",
        platform_unknown: "Desconhecido",
        prog_processing: "Processando...",
        prog_done: "Concluído!",
        prog_err: "Erro"
    },
    en: {
        nav_convert_label: "Convert",
        nav_convert_btn: "Convert ISO to GOD",
        nav_copy_btn: "Copy DVD to ISO",
        nav_about_btn: "About",
        nav_support_btn: "Support the Project",
        status_checking: "Checking...",
        status_deps_ok: "Dependencies OK",
        status_missing: "Missing: ",
        status_err_verify: "Error checking",
        page_convert_title: "ISO → GOD",
        page_convert_desc: "Transform ISO files into GOD format for use on Xbox 360 (RGH/JTAG)",
        label_iso_input: "Input ISO File",
        placeholder_select_iso: "Select ISO file...",
        btn_browse: "Browse",
        label_god_output: "GOD Output Directory",
        placeholder_select_dir: "Select directory...",
        btn_convert_action: "Convert to GOD",
        progress_waiting: "Waiting...",
        page_copy_title: "Copy ISO from DVD",
        page_copy_desc: "Create an ISO image from your Xbox 360 game DVD",
        label_dvd_drive: "DVD Drive",
        select_drive_default: "Select a drive...",
        select_drive_searching: "Searching drives...",
        select_drive_empty: "No optical drive found",
        select_drive_err: "Error listing drives",
        label_output_dir: "Output Directory",
        label_iso_name: "ISO File Name",
        placeholder_iso_name: "my_game.iso",
        btn_copy_action: "Start Copy",
        about_title: "About XboxForGOD",
        about_desc: "How it works and what it is for",
        about_q1: "What is GOD format?",
        about_a1: "The <strong>GOD (Games on Demand)</strong> format is a specific file format for Xbox 360 games, designed to be executed directly from the console's hard drive. It is widely used in modified Xbox 360 consoles with <strong>RGH/JTAG</strong>.",
        about_q2: "How does it work?",
        about_step1_title: "Insert DVD",
        about_step1_desc: "Insert the Xbox 360 game DVD into the computer's optical drive.",
        about_step2_title: "Copy ISO",
        about_step2_desc: "The program uses the <code>dd</code> command (Linux) or direct reading (Windows) to create an identical ISO image of the disc.",
        about_step3_title: "Convert to GOD",
        about_step3_desc: "The ISO is converted to GOD format using the <code>iso2god</code> utility, generating the folder structure compatible with Xbox 360.",
        about_step4_title: "Copy to HDD",
        about_step4_desc: "The generated GOD folder is copied to the internal or external HDD of the Xbox 360 with exploit (RGH/JTAG), appearing directly in the games menu.",
        about_adv_title: "Advantages of GOD format",
        about_adv1: "<strong>Convenience</strong> — Games run directly from HDD, without physical disc.",
        about_adv2: "<strong>Speed</strong> — Faster loading than DVD reading.",
        about_adv3: "<strong>Preservation</strong> — Reduces wear on the console's DVD reader.",
        about_adv4: "<strong>Organization</strong> — Facilitates the management of large game libraries.",
        about_req_title: "System Requirements",
        about_req_detected: "Detected Platform",
        about_dev_by: "Developed by",
        
        log_err_dir: "Error selecting directory: ",
        log_err_file: "Error selecting file: ",
        log_warn_dvd: "Select a DVD drive.",
        log_warn_dir: "Select an output directory.",
        log_warn_name: "Enter a name for the ISO file.",
        log_warn_iso: "Select an ISO file.",
        btn_copying: "Copying...",
        btn_converting: "Converting...",
        log_done: "Done: ",
        log_error: "Error: ",
        platform_unknown: "Unknown",
        prog_processing: "Processing...",
        prog_done: "Done!",
        prog_err: "Error"
    }
};

let currentLang = localStorage.getItem('xboxforgod_lang') || 'pt';

function t(key) {
    return translations[currentLang][key] || key;
}

function applyLanguage() {
    document.querySelectorAll('[data-i18n]').forEach(el => {
        const key = el.getAttribute('data-i18n');
        const isHtml = el.getAttribute('data-i18n-html');
        if (isHtml === 'true') {
            el.innerHTML = t(key);
        } else {
            el.textContent = t(key);
        }
    });

    document.querySelectorAll('[data-i18n-placeholder]').forEach(el => {
        const key = el.getAttribute('data-i18n-placeholder');
        el.placeholder = t(key);
    });
}

function setupLanguageSelector() {
    const btns = document.querySelectorAll('.lang-btn');
    
    btns.forEach(btn => {
        if (btn.dataset.lang === currentLang) {
            btn.classList.add('active');
        } else {
            btn.classList.remove('active');
        }

        btn.addEventListener('click', () => {
            currentLang = btn.dataset.lang;
            localStorage.setItem('xboxforgod_lang', currentLang);
            
            btns.forEach(b => b.classList.remove('active'));
            btn.classList.add('active');
            
            applyLanguage();
            checkDependencies();
            
            const driveSelect = document.getElementById('drive-select');
            if (driveSelect.options.length === 1 && driveSelect.options[0].value === "") {
               refreshDrives();
            } else if (driveSelect.options.length > 0 && driveSelect.options[0].value === "") {
                driveSelect.options[0].textContent = t('select_drive_default');
            }
        });
    });
}

document.addEventListener('DOMContentLoaded', init);

function init() {
    applyLanguage();
    setupLanguageSelector();
    setupNavigation();
    setupCopyISO();
    setupConvertGOD();
    checkDependencies();
    loadPlatformInfo();
    refreshDrives();

    const supportBtn = document.getElementById('nav-support');
    if (supportBtn) {
        supportBtn.addEventListener('click', () => {
            BrowserOpenURL('https://www.paypal.com/ncp/payment/8V6WQCGN6HDCQ');
        });
    }
}

function setupNavigation() {
    const navButtons = document.querySelectorAll('.nav-btn');
    const pages = document.querySelectorAll('.page');

    navButtons.forEach(btn => {
        btn.addEventListener('click', () => {
            const targetPage = btn.dataset.page;

            navButtons.forEach(b => b.classList.remove('active'));
            btn.classList.add('active');

            pages.forEach(p => {
                p.classList.remove('active');
                if (p.id === 'page-' + targetPage) {
                    p.classList.add('active');
                }
            });
        });
    });
}

function setupCopyISO() {
    document.getElementById('btn-refresh-drives').addEventListener('click', refreshDrives);

    document.getElementById('btn-select-output-dir').addEventListener('click', async () => {
        try {
            const dir = await SelecionarDiretorio();
            if (dir) {
                document.getElementById('iso-output-dir').value = dir;
            }
        } catch (err) {
            appendLog('copy-log', t('log_err_dir') + err, true);
        }
    });

    document.getElementById('btn-copy-iso').addEventListener('click', async () => {
        const driveSelect = document.getElementById('drive-select');
        const outputDir = document.getElementById('iso-output-dir').value;
        const outputName = document.getElementById('iso-output-name').value;

        if (!driveSelect.value) {
            appendLog('copy-log', t('log_warn_dvd'), true);
            return;
        }
        if (!outputDir) {
            appendLog('copy-log', t('log_warn_dir'), true);
            return;
        }
        if (!outputName) {
            appendLog('copy-log', t('log_warn_name'), true);
            return;
        }

        const btn = document.getElementById('btn-copy-iso');
        btn.disabled = true;
        const btnText = btn.querySelector('span');
        btnText.textContent = t('btn_copying');
        
        showLog('copy-progress-panel');
        clearLog('copy-log');

        try {
            const result = await CopiarISO(driveSelect.value, outputDir, outputName);
            appendLog('copy-log', t('log_done') + result);
        } catch (err) {
            appendLog('copy-log', t('log_error') + err, true);
        } finally {
            btn.disabled = false;
            btnText.textContent = t('btn_copy_action');
        }
    });
}

function setupConvertGOD() {
    document.getElementById('btn-select-iso').addEventListener('click', async () => {
        try {
            const file = await SelecionarArquivoISO();
            if (file) {
                document.getElementById('convert-iso-path').value = file;
            }
        } catch (err) {
            appendLog('convert-log', t('log_err_file') + err, true);
        }
    });

    document.getElementById('btn-select-god-dir').addEventListener('click', async () => {
        try {
            const dir = await SelecionarDiretorio();
            if (dir) {
                document.getElementById('convert-output-dir').value = dir;
            }
        } catch (err) {
            appendLog('convert-log', t('log_err_dir') + err, true);
        }
    });

    document.getElementById('btn-convert-god').addEventListener('click', async () => {
        const isoPath = document.getElementById('convert-iso-path').value;
        const outputDir = document.getElementById('convert-output-dir').value;

        if (!isoPath) {
            appendLog('convert-log', t('log_warn_iso'), true);
            return;
        }

        const btn = document.getElementById('btn-convert-god');
        btn.disabled = true;
        const btnText = btn.querySelector('span');
        btnText.textContent = t('btn_converting');
        
        showLog('convert-progress-panel');
        clearLog('convert-log');

        try {
            const result = await ConverterISOparaGOD(isoPath, outputDir);
            appendLog('convert-log', result);
        } catch (err) {
            appendLog('convert-log', t('log_error') + err, true);
        } finally {
            btn.disabled = false;
            btnText.textContent = t('btn_convert_action');
        }
    });

    EventsOn('progress', (ev) => {
        const copyPanel = document.getElementById('copy-progress-panel');
        const convertPanel = document.getElementById('convert-progress-panel');

        if (copyPanel && copyPanel.style.display === 'block') {
            updateProgress('copy-progress', ev);
        }
        if (convertPanel && convertPanel.style.display === 'block') {
            updateProgress('convert-progress', ev);
        }
    });
}

async function checkDependencies() {
    const dot = document.querySelector('.status-dot');
    const text = document.getElementById('dep-status-text');

    try {
        const status = await VerificarDependencias();
        if (status.ddInstalled && status.iso2godInstalled) {
            dot.className = 'status-dot ok';
            text.textContent = t('status_deps_ok');
        } else {
            dot.className = 'status-dot error';
            const missing = [];
            if (!status.ddInstalled) missing.push('dd');
            if (!status.iso2godInstalled) missing.push('iso2god');
            text.textContent = t('status_missing') + missing.join(', ');
        }
    } catch (err) {
        dot.className = 'status-dot error';
        text.textContent = t('status_err_verify');
    }
}

async function loadPlatformInfo() {
    try {
        const platform = await GetPlatform();
        const label = platform === 'linux' ? 'Linux' : platform === 'windows' ? 'Windows' : platform;
        document.getElementById('platform-info').textContent = label;
    } catch (_) {
        document.getElementById('platform-info').textContent = t('platform_unknown');
    }
}

async function refreshDrives() {
    const select = document.getElementById('drive-select');
    select.innerHTML = `<option value="">${t('select_drive_searching')}</option>`;

    try {
        const drives = await ListarDispositivosCDROM();
        select.innerHTML = `<option value="">${t('select_drive_default')}</option>`;

        if (drives && drives.length > 0) {
            drives.forEach(d => {
                const opt = document.createElement('option');
                opt.value = d.path;
                opt.textContent = d.name + ' (' + d.path + ') — ' + (d.size || 'N/A');
                select.appendChild(opt);
            });
        } else {
            select.innerHTML = `<option value="">${t('select_drive_empty')}</option>`;
        }
    } catch (err) {
        select.innerHTML = `<option value="">${t('select_drive_err')}</option>`;
    }
}

function showLog(panelId) {
    document.getElementById(panelId).style.display = 'block';
}

function clearLog(logId) {
    document.getElementById(logId).textContent = '';
}

function appendLog(logId, message, isError) {
    const el = document.getElementById(logId);
    if (!el) return;
    const panel = el.closest('.progress-panel');
    if (panel) panel.style.display = 'block';

    const line = document.createElement('div');
    line.textContent = message;
    if (isError) line.style.color = '#ef4444';
    el.appendChild(line);
    el.scrollTop = el.scrollHeight;
}

function updateProgress(prefix, ev) {
    const title = document.getElementById(prefix + '-title');
    const percent = document.getElementById(prefix + '-percent');
    const fill = document.getElementById(prefix + '-fill');
    const glow = document.getElementById(prefix + '-glow');
    const details = document.getElementById(prefix + '-details');
    
    if (ev.type === 'start') {
        title.textContent = ev.message;
        percent.textContent = '0%';
        fill.style.width = '0%';
        glow.style.width = '0%';
        fill.style.background = '';
        glow.style.background = '';
        details.textContent = '';
    } else if (ev.type === 'progress') {
        title.textContent = t('prog_processing');
        percent.textContent = ev.percent.toFixed(1) + '%';
        fill.style.width = ev.percent + '%';
        glow.style.width = ev.percent + '%';
        details.textContent = ev.message;
    } else if (ev.type === 'complete') {
        title.textContent = t('prog_done');
        percent.textContent = '100%';
        fill.style.width = '100%';
        glow.style.width = '100%';
        details.textContent = ev.message;
    } else if (ev.type === 'error') {
        title.textContent = t('prog_err');
        details.textContent = ev.message;
        fill.style.background = 'var(--danger)';
        glow.style.background = 'var(--danger)';
        appendLog(prefix.replace('-progress', '-log'), ev.message, true);
    } else if (ev.type === 'log') {
        appendLog(prefix.replace('-progress', '-log'), ev.message);
    }
}
